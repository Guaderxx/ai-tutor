package runpy

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"
)

func TextToAudio(text string) (audioPath string) {
	command := "python"
	audioPath = fmt.Sprintf("tmp/%d.wav", time.Now().Unix())
	args := []string{"pyfile/long-text-to-audio.py", text, audioPath}

	cmd := exec.Command(command, args...)
	// TODO
	cmd.Stdout = ioutil.Discard
	cmd.Stderr = ioutil.Discard

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	return
}
