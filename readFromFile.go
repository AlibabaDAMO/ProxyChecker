package main

import (
	"bufio"
	"os"
)

//Reading proxies from file
func readFromFile(path string, proxyType int) []string {

	var proxies []string

	file, err := os.Open(path)

	if err != nil {

		println("Can't open file")
	}

	defer file.Close()

	fileScaner := bufio.NewScanner(file)

	for fileScaner.Scan() {

		//Appending proxies to slice
		switch proxyType {
		case 0:
			proxies = append(proxies, "http://"+fileScaner.Text())
		case 1:
			proxies = append(proxies, fileScaner.Text())
		}
	}

	return proxies
}
