package code

import (
	"golang.org/x/net/proxy"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"time"
)

const (
	timeout = time.Duration(5 * time.Second)
)

//CheckProxySOCKS Check proxies on valid
func CheckProxySOCKS(proxyy string, c chan QR, wg *sync.WaitGroup) (err error) {

	defer wg.Done()

	d := net.Dialer{
		Timeout:   timeout,
		KeepAlive: timeout,
	}

	//Sending request through proxy
	dialer, _ := proxy.SOCKS5("tcp", proxyy, nil, &d)

	httpClient := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			DisableKeepAlives: true,
			Dial:              dialer.Dial,
		},
	}
	response, err := httpClient.Get("https://api.ipify.org?format=json")

	if err != nil {

		c <- QR{Addr: proxyy, Res: false}
		return
	}

	defer response.Body.Close()
	io.Copy(ioutil.Discard, response.Body)

	c <- QR{Addr: proxyy, Res: true}

	return nil

}
