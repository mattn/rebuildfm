// +build !darwin

package main

import (
	"errors"
	"os"
	"os/exec"
)

func playURL(url string) error {
	for _, player := range []string{"ffplay", "avplay"} {
		f, err := exec.LookPath(player)
		if err == nil {
			args := []string{"-autoexit", "-nodisp", url}
			return exec.Command(f, args...).Run()
		}
	}

	f, err := exec.LookPath("mplayer")
	if err == nil {
		cmd := exec.Command(f, url)
		cmd.Stdin = os.Stdin
		return cmd.Run()
	}
	return errors.New("player not found")
}

