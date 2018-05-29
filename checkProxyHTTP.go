package main

import (
	"github.com/trigun117/ProxyChecker/code"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"time"
)

//CheckProxyHTTP Check proxies on valid
func checkProxyHTTP(proxy string, c chan code.QR, realIP string, wg *sync.WaitGroup) (err error) {

	defer wg.Done()

	//Sending request through proxy
	proxyURL, _ := url.Parse(proxy)

	timeout := time.Duration(5 * time.Second)
	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Proxy:             http.ProxyURL(proxyURL),
		},
	}
	response, err := httpClient.Get("https://api.ipify.org?format=json")

	if err != nil {

		c <- code.QR{Addr: proxy, Res: false}
		return
	}

	defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)

	resp := code.CleanIP(string(body))

	//Proxy is checked for anonymity
	if resp != realIP {

		c <- code.QR{Addr: proxy, Res: true}
		return
	}

	return nil
}
