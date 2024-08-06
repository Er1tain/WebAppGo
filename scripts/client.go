package scripts

import (
	"log"
	"os"
	"os/exec"
)

func RunClient() {
	cmdReact := exec.Command("npm", "start")
	cmdReact.Dir = "./view"
	cmdReact.Stdout = os.Stdout
	cmdReact.Stderr = os.Stderr

	errReact := cmdReact.Start()
	if errReact != nil {
		log.Fatal(errReact)
	}

	errReactWait := cmdReact.Wait()
	if errReactWait != nil {
		log.Fatal(errReactWait)
	}
}
