package main

import (
	"log"
	"os/exec"
	"regexp"
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

func TruncateString(org string) string {
	const maxLength = 40
	re := regexp.MustCompile(`\r?\n\s+`)
	ellipsed := re.ReplaceAllString(org, " ")
	if len(ellipsed) > maxLength {
		ellipsed = org[0:37] + "..."
	}

	return ellipsed
}

func GetPassedTime(date string) string {
	log.Println(date)
	/*
		t, err := time.Parse(date)
		if err != nil {
			panic(err)
		}
	*/

}
