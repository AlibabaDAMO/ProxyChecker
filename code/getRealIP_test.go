package code

import (
	"testing"
)

func TestGetRealIP(t *testing.T) {

	if str, err := GetRealIP(`https://api.ipify.org?format=json`); str == "" && err != nil {
		t.Fail()
	}

	if _, err := GetRealIP(`test`); err == nil {
		t.Fail()
	}
}
