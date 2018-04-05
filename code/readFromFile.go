package code

import (
	"bufio"
	"os"
)

//ReadFromFile Reading proxies from file
func ReadFromFile(path string, proxyType int) ([]string, error) {

	var proxies []string

	file, err := os.Open(path)
	if err != nil {
		return proxies, err
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

	return proxies, err
}
