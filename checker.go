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

func writeToFile(proxyURL string) {

	fmt.Println("Live proxy is: ", proxyURL)

	file, _ := os.OpenFile(`live-proxy.txt`, os.O_APPEND, 0666)

	defer file.Close()

	fileWriter := bufio.NewWriter(file)

	fmt.Fprintln(fileWriter, proxyURL)

	fileWriter.Flush()
}

func checkProxy(proxy string, c chan QR) {

	fmt.Println("Get proxy", proxy)

	proxyURL, _ := url.Parse(proxy)
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout, Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	_, err := httpClient.Get("https://www.google.com/")

	if err != nil {

		fmt.Println("Dead proxy", proxy)
		c <- QR{Addr: proxy, Res: false}
	} else {

		c <- QR{Addr: proxy, Res: true}
	}
}

func readFromFile(path string) []string {

	var proxys []string

	file, err := os.Open(path)

	if err != nil {

		fmt.Println("Can't open file")
	}

	defer file.Close()

	fileScaner := bufio.NewScanner(file)

	for fileScaner.Scan() {

		proxys = append(proxys, "http://"+fileScaner.Text())
	}

	fmt.Println("Got ", len(proxys), " proxys")

	return proxys
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	respChan := make(chan QR)

	fmt.Printf("Enter path to file: ")
	var path string
	fmt.Scan(&path)
	os.Create(`live-proxy.txt`)
	prox := readFromFile(path)

	for _, proxy := range prox {

		go checkProxy(proxy, respChan)
	}

	for range prox {
		r := <-respChan

		if r.Res {
			writeToFile(r.Addr)
		}
	}
}
