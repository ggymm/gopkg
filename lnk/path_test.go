package lnk

import (
	"bufio"
	"os"
	"testing"
)

func Test_ParsePath(t *testing.T) {
	file, err := os.Open("C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\Google Chrome.lnk")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = file.Close()
	}()
	lnk, err := ParsePath(bufio.NewReader(file))
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", lnk)
}
