package code

import (
	"testing"
)

func TestWriteToFile(t *testing.T) {
	if WriteToFile(`000.00.000.000:0000`) != nil {
		t.Fail()
	}
}
