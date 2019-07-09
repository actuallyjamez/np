package main

import (
	"log"
	"os"
	"os/exec"
)

func clearArg(env string) {
	err := os.Unsetenv(env)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		os.Exit(0)
	}

	clearArg("http_proxy")
	clearArg("https_proxy")
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err := cmd.Run()
	if err != nil {
		//log.Fatal(err)
		os.Exit(1)
	}
}
