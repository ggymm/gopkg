package files

import (
	"os"
	"testing"
)

func TestCurrentPath(t *testing.T) {
	t.Log(CurrentPath())
}

func TestMkdir(t *testing.T) {
	err := Mkdir("D:\\temp\\test")
	if err != nil {
		t.Errorf("Mkdir() error = %v", err)
		return
	}
	t.Log("success")
	_ = os.RemoveAll("D:\\temp\\test")
}
