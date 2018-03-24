package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"time"
)

//QR contain info about proxy
type QR struct {
	Addr string
	Res  bool
}

//Removes everything around ip
func cleanIP(dirtyIP string) (cleanIP string) {

	r1, _ := regexp.Compile(`^.......|..$`)
	cleanIP = r1.ReplaceAllString(dirtyIP, "")

	return
}

//Getting real ip of device
func getRealIP() (realIP string) {

	res, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {

		panic(err)
	}

	body1, _ := ioutil.ReadAll(res.Body)

	realIP = cleanIP(string(body1))

	return
}

//Check proxies on uniqueness
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

//Check proxies on valid
func checkProxy(proxy string, c chan QR, realIP string) {

	//Sending request through proxy
	proxyURL, _ := url.Parse(proxy)
	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{Timeout: timeout, Transport: &http.Transport{Proxy: http.ProxyURL(proxyURL)}}
	response, err := httpClient.Get("https://api.ipify.org?format=json")

	if err != nil {

		c <- QR{Addr: proxy, Res: false}
	} else {

		body, _ := ioutil.ReadAll(response.Body)
		resp := cleanIP(string(body))

		//Proxy is checked for anonymity
		if resp != realIP {

			c <- QR{Addr: proxy, Res: true}
		}
	}
}

//Reading proxies from file
func readFromFile(path string) []string {

	var proxies []string

	file, err := os.Open(path)

	if err != nil {

		fmt.Println("Can't open file")
	}

	defer file.Close()

	fileScaner := bufio.NewScanner(file)

	for fileScaner.Scan() {

		//Appending proxies to slice
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

	//Creating file live-proxies.txt
	os.Create(`live-proxies.txt`)

	prox := readFromFile(path)

	uniqueProxies := unique(prox)

	fmt.Println("Got", len(uniqueProxies), "unique proxies")

	time.Sleep(2 * time.Second)

	fmt.Println("START")

	realIP := getRealIP()

	for _, proxy := range uniqueProxies {

		go checkProxy(proxy, respChan, realIP)
	}

	for range uniqueProxies {
		r := <-respChan

		if r.Res {
			writeToFile(r.Addr)
		}
	}
}
