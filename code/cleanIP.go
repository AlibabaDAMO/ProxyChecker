package code

import (
	"regexp"
)

//CleanIP Removes everything around ip
func CleanIP(dirtyIP string) (cleanIP string) {

	r1, _ := regexp.Compile(`^.......|..$`)
	cleanIP = r1.ReplaceAllString(dirtyIP, "")

	return
}
