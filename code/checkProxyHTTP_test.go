package code

import (
	"testing"
)

func TestCheckProxyHTTP(t *testing.T) {

	ch := make(chan QR, 1)

	realIP, _ := GetRealIP(`https://api.ipify.org?format=json`)

	CheckProxyHTTP(`http://000.00.000.000:0000`, ch, realIP)

	r := <-ch

	if r.Res != false {
		t.Fail()
	}

	CheckProxyHTTP(`http://80.211.141.30:8888`, ch, realIP)

	r = <-ch

	if r.Res != true {
		t.Fail()
	}
}
