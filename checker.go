package main

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"time"
)

//QR contain info about proxy
type QR struct {
	Addr string
	Res  bool
}

func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func writeToFile(proxyURL string) {

	file, _ := os.OpenFile(`live-proxies.txt`, os.O_APPEND, 0666)

	defer file.Close()

	fileWriter := bufio.NewWriter(file)

	fmt.Fprintln(fileWriter, proxyURL)

	fileWriter.Flush()
}

func checkProxy(proxy string, c chan QR) {

	proxyURL, _ := url.Parse(proxy)
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout, Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	_, err := httpClient.Get("https://www.google.com/")

	if err != nil {

		c <- QR{Addr: proxy, Res: false}
	} else {

		c <- QR{Addr: proxy, Res: true}
	}
}

func readFromFile(path string) []string {

	var proxies []string

	file, err := os.Open(path)

	if err != nil {

		fmt.Println("Can't open file")
	}

	defer file.Close()

	fileScaner := bufio.NewScanner(file)

	for fileScaner.Scan() {

		proxies = append(proxies, "http://"+fileScaner.Text())
	}

	fmt.Println("Got", len(proxies), "proxies from file")

	return proxies
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	respChan := make(chan QR)

	fmt.Printf("Enter path to file: ")
	var path string
	fmt.Scan(&path)

	os.Create(`live-proxies.txt`)

	prox := readFromFile(path)

	uniqueProxies := unique(prox)

	fmt.Println("Got", len(uniqueProxies), "unique proxies")

	time.Sleep(5 * time.Second)

	fmt.Println("START")

	for _, proxy := range uniqueProxies {

		go checkProxy(proxy, respChan)
	}

	for range uniqueProxies {
		r := <-respChan

		if r.Res {
			writeToFile(r.Addr)
		}
	}
}
