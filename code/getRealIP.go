package code

import (
	"io/ioutil"
	"net/http"
)

//GetRealIP Getting real ip of device
func GetRealIP(link string) (string, error) {

	res, err := http.Get(link)
	if err != nil {

		return "", err
	}

	body1, _ := ioutil.ReadAll(res.Body)

	realIP := cleanIP(string(body1))

	return realIP, err
}
