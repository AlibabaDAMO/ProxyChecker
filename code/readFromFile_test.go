package code

import (
	"os"
	"testing"
)

func TestReadFromFile(t *testing.T) {

	if _, err := ReadFromFile(`test@!#$%`, 1); err == nil {
		t.Fail()
	}

	os.Create(`TestReadFromFile.txt`)

	f, _ := os.OpenFile(`TestReadFromFile.txt`, os.O_APPEND, 0644)

	f.WriteString(`000.00.000.000:0000`)

	if _, err := ReadFromFile(`TestReadFromFile.txt`, 1); err != nil {
		t.Fail()
	}

	if _, err := ReadFromFile(`TestReadFromFile.txt`, 0); err != nil {
		t.Fail()
	}
}
