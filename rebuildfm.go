package main

import (
	"bytes"
	"code.google.com/p/go.net/html"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"runtime"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Items   Items    `xml:"channel"`
}

type Items struct {
	XMLName  xml.Name `xml:"channel"`
	ItemList []Item   `xml:"item"`
}

type Item struct {
	Title       string      `xml:"title"`
	Link        string      `xml:"link"`
	Description string      `xml:"description"`
	Enclosure   []Enclosure `xml:"enclosure"`
}

type Enclosure struct {
	URL string `xml:"url,attr"`
}

var blockTags = []string {"div", "br", "p", "blockquote", "pre", "h1", "h2", "h3", "h4", "h5", "h6"}

func extractText(node *html.Node, w *bytes.Buffer) {
	if node.Type == html.TextNode {
		data := strings.Trim(node.Data, "\r\n")
		if data != "" {
			w.WriteString(data)
		}
	} else if node.Type == html.ElementNode {
		if node.Data == "li" {
			w.WriteString("\n* ")
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		extractText(c, w)
	}
	if node.Type == html.ElementNode {
		for _, bt := range blockTags {
			if strings.ToLower(node.Data) == bt {
				w.WriteString("\n")
				break
			}
		}
	}
}

func play(items ...Item) error {
	for _, i := range items {
		doc, err := html.Parse(strings.NewReader(i.Description))
		if err != nil {
			return err
		}
		var buf bytes.Buffer
		extractText(doc, &buf)

		fmt.Println(i.Title + "\n")
		fmt.Println(buf.String())
		for _, e := range i.Enclosure {
			args := []string{"-autoexit", "-nodisp", e.URL}

			if runtime.GOOS == "darwin" {
				if err :=  _play(e.URL); err != nil {
					return err
				}
			} else {
				if err := exec.Command("ffplay", args...).Run(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func main() {
	flag.Parse()

	res, err := http.Get("http://feeds.rebuild.fm/rebuildfm")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	var rss RSS
	err = xml.NewDecoder(res.Body).Decode(&rss)
	if err != nil {
		log.Fatal(err)
	}

	if flag.NArg() == 0 {
		for _, i := range rss.Items.ItemList {
			fmt.Println(i.Title)
		}
		return
	}

	ep := flag.Arg(0)
	if ep != "-" {
		for _, i := range rss.Items.ItemList {
			if strings.HasPrefix(i.Title, ep+":") {
				err = play(i)
				if err != nil {
					log.Fatal(err)
				}
				return
			}
		}
		log.Fatal("404 Episode Not Found")
	} else {
		err = play(rss.Items.ItemList...)
		if err != nil {
			log.Fatal(err)
		}
	}
}
