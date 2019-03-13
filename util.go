package main

import (
	"os/exec"
	"runtime"
)

func OpenBrowser(url string) bool {
	var args string
	switch runtime.GOOS {
	case "darwin":
		args = "open"
	default:
		args = "xdg-open"
	}

	cmd := exec.Command(args, url)
	return cmd.Start() == nil
}
