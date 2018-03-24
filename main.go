package main

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	respChan := make(chan QR)

	fmt.Printf("Enter path to file: ")
	var path string
	fmt.Scan(&path)

	prox := readFromFile(path)

	uniqueProxies := unique(prox)

	fmt.Println("Got", len(uniqueProxies), "unique proxies")

	time.Sleep(2 * time.Second)

	fmt.Println("START")

	realIP := getRealIP()

	for _, proxy := range uniqueProxies {

		go checkProxy(proxy, respChan, realIP)
	}

	os.Create(`live-proxies.txt`)

	for range uniqueProxies {
		r := <-respChan

		if r.Res {
			writeToFile(r.Addr)
		}
	}
}
