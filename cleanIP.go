package main

import (
	"regexp"
)

//Removes everything around ip
func cleanIP(dirtyIP string) (cleanIP string) {

	r1, _ := regexp.Compile(`^.......|..$`)
	cleanIP = r1.ReplaceAllString(dirtyIP, "")

	return
}
