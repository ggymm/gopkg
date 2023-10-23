package files

import (
	"testing"
)

func TestReadFileByLine(t *testing.T) {
	content, err := ReadLines("read_test.go")
	if err != nil {
		t.Errorf("ReadLines() error = %v", err)
		return
	}
	t.Log(content)
}

func TestReadFileToString(t *testing.T) {
	content, err := ReadString("read_test.go")
	if err != nil {
		t.Errorf("ReadString() error = %v", err)
		return
	}
	t.Log(content)
}

func TestReadFileToBytes(t *testing.T) {
	content, err := ReadBytes("read_test.go")
	if err != nil {
		t.Errorf("ReadBytes() error = %v", err)
		return
	}
	t.Log(content)
}
