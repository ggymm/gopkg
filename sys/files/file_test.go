package files

import "testing"

func TestCopyFile(t *testing.T) {
	var (
		src = ""
		dst = ""
	)

	err := CopyFile(src, dst)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("CopyFile success")
}

func TestExists(t *testing.T) {
	exists := Exist("file.go")
	if exists {
		t.Log("file exists")
	} else {
		t.Log("file not exists")
	}
}
