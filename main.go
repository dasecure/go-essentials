package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	file, err := os.Open(pidFile)
	if err != nil {
		return err
	}
	defer file.Close()

	var pid int
	if _, err := fmt.Fscanf(file, "%d", &pid); err != nil {
		return errors.Wrap(err, "failed to read pid")
	}
	fmt.Printf("kill server with pid: %d", pid)

	if err := os.Remove(pidFile); err != nil {
		log.Printf("failed to remove pid file: %v", err)
	}
	return nil
}

func main() {
	if err := killServer("server.pid"); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(1)
	}
}
