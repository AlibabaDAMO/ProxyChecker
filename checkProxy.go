package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

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
