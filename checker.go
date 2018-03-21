package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

func writeToFile(proxyURL string) {

	fmt.Println("Live proxy is: ", proxyURL)

	file, _ := os.OpenFile(`live-proxy.txt`, os.O_APPEND, 0666)

	defer file.Close()

	fileWriter := bufio.NewWriter(file)

	fmt.Fprintln(fileWriter, proxyURL)

	fileWriter.Flush()
}

func checkProxy(proxy string) (deadproxy string, err error) {

	fmt.Println("Get proxy", proxy)

	proxyURL, _ := url.Parse(proxy)
	timeout := time.Duration(1 * time.Second)
	httpClient := &http.Client{Timeout: timeout, Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	re, err := httpClient.Get("https://www.google.com/")

	if err != nil {

		fmt.Println("Dead proxy", proxy)
		return
	}

	writeToFile(proxy)

	re.Body.Close()

	return
}

func readFromFile(path string) {

	file, err := os.Open(path)

	if err != nil {

		fmt.Println("Can't open file")
	}

	defer file.Close()

	fileScaner := bufio.NewScanner(file)

	for fileScaner.Scan() {

		if checkProxy("http://" + fileScaner.Text()); err == nil {

			continue
		}
	}
}

func main() {

	fmt.Printf("Enter path to file: ")
	var path string
	fmt.Scan(&path)
	os.Create(`live-proxy.txt`)
	readFromFile(path)
}
