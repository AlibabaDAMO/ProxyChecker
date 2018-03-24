package main

import (
	"bufio"
	"os"
)

//Reading proxies from file
func readFromFile(path string) []string {

	var proxies []string

	file, err := os.Open(path)

	if err != nil {

		println("Can't open file")
	}

	defer file.Close()

	fileScaner := bufio.NewScanner(file)

	for fileScaner.Scan() {

		//Appending proxies to slice
		proxies = append(proxies, "http://"+fileScaner.Text())
	}

	println("Got", len(proxies), "proxies from file")

	return proxies
}
