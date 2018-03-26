package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

//Writing valid proxies to file
func writeToFile(proxyURL string) {

	//Opening file live-proxies.txt
	file, _ := os.OpenFile(`live-proxies.txt`, os.O_APPEND, 0666)

	defer file.Close()

	fileWriter := bufio.NewWriter(file)

	//Remove http from proxy url
	r, _ := regexp.Compile(`^http://`)

	cleanProxy := r.ReplaceAllString(proxyURL, "")

	//Writing to file
	fmt.Fprintln(fileWriter, cleanProxy)

	fileWriter.Flush()
}
