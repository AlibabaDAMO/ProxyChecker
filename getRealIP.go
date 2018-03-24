package main

import (
	"io/ioutil"
	"net/http"
)

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
