package main

import (
	"bufio"
	"fmt"
	"os"
)

//Writing valid proxies to file
func writeToFile(proxyURL string) {

	//Opening file live-proxies.txt
	file, _ := os.OpenFile(`live-proxies.txt`, os.O_APPEND, 0666)

	defer file.Close()

	fileWriter := bufio.NewWriter(file)

	//Writing to file
	fmt.Fprintln(fileWriter, proxyURL)

	fileWriter.Flush()
}
