package code

import (
	"os"
	"regexp"
)

//WriteToFile Writing valid proxies to file
func WriteToFile(proxyURL string) error {

	//Opening file live-proxies.txt
	file, _ := os.OpenFile("live-proxies.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer file.Close()

	r := regexp.MustCompile(`^http://`)

	cleanProxy := r.ReplaceAllString(proxyURL, "")

	_, err := file.WriteString(cleanProxy + "\r\n")

	return err
}
