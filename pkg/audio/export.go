package audio

import (
	"io"
	"log"
	"os/exec"
)

func PlayAudio(audioPath string) {
	command := "cvlc"
	args := []string{"--play-and-exit", audioPath}

	cmd := exec.Command(command, args...)
	cmd.Stdout = io.Discard
	cmd.Stderr = io.Discard

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
