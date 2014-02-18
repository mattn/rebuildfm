// +build !darwin
package main

import (
	"os/exec"
)

func playURL(url string) error {
	args := []string{"-autoexit", "-nodisp", url}
	return exec.Command("ffplay", args...).Run()
}

