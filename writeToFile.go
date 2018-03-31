package main

import (
	"os"
	"regexp"
)

//Writing valid proxies to file
func writeToFile(proxyURL string) {

	//Opening file live-proxies.txt
	file, _ := os.OpenFile("live-proxies.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	r := regexp.MustCompile(`^http://`)

	cleanProxy := r.ReplaceAllString(proxyURL, "")

	file.WriteString(cleanProxy + "\r\n")
}
