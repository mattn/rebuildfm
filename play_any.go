// +build !darwin

package main

import (
	"errors"
	"os/exec"
)

func playURL(url string) error {
	f, err := exec.LookPath("ffplay")
	if err == nil {
		args := []string{"-autoexit", "-nodisp", url}
		return exec.Command(f, args...).Run()
	}
	f, err = exec.LookPath("mplayer")
	if err == nil {
		return exec.Command(f, url).Run()
	}
	return errors.New("player not found")
}

