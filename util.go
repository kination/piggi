package main

import (
	"os/exec"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
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
	t, err := time.Parse(time.RFC3339, date)
	var passedTime string
	if err != nil {
		panic(err)
	}

	if time.Now().Sub(t).Seconds() < 60 {
		passedTime = strings.Join(
			[]string{strconv.Itoa(int(time.Now().Sub(t).Seconds())), " seconds before"}, "")
	} else if time.Now().Sub(t).Minutes() < 60 {
		passedTime = strings.Join(
			[]string{strconv.Itoa(int(time.Now().Sub(t).Minutes())), " minutes before"}, "")
	} else if time.Now().Sub(t).Hours() < 60 {
		passedTime = strings.Join(
			[]string{strconv.Itoa(int(time.Now().Sub(t).Hours())), " hours before"}, "")
	} else {
		days := int(time.Now().Sub(t).Hours()) / 24
		passedTime = strings.Join(
			[]string{strconv.Itoa(days), " days before"}, "")
	}

	return passedTime
}
