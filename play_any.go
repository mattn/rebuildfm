// +build !darwin

package main

import (
	"errors"
	"os"
	"os/exec"
)

func playURL(url string) error {
	var cmd *exec.Cmd
	for _, player := range []string{"ffplay", "avplay"} {
		f, err := exec.LookPath(player)
		if err == nil {
			args := []string{"-autoexit", "-nodisp", url}
			cmd = exec.Command(f, args...)
			cmd.Stdin = os.Stdin
			return cmd.Run()
		}
	}

	f, err := exec.LookPath("mplayer")
	if err == nil {
		cmd = exec.Command(f, url)
		cmd.Stdin = os.Stdin
		return cmd.Run()
	}
	return errors.New("player not found")
}

