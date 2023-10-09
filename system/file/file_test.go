package file

import (
	"mime/multipart"
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

func TestReadFileByLine(t *testing.T) {
	content, err := ReadFileToLines("file_test.go")
	if err != nil {
		t.Errorf("ReadFileToLines() error = %v", err)
		return
	}
	t.Log(content)
}

func TestReadFileToString(t *testing.T) {
	content, err := ReadFileToString("file_test.go")
	if err != nil {
		t.Errorf("ReadFileToString() error = %v", err)
		return
	}
	t.Log(content)
}

func TestReadMultipartFileToBytes(t *testing.T) {
	// 创建 multipart.FileHeader
	fh := &multipart.FileHeader{
		Filename: "file_test.go",
		// TODO: 构建 multipart.FileHeader
	}
	content, err := ReadMultipartFileToBytes(fh)
	if err != nil {
		t.Errorf("ReadMultipartFileToBytes() error = %v", err)
		return
	}
	t.Log(content)
}

func TestWriteBytesToFile(t *testing.T) {
	err := WriteBytesToFile("D:\\temp\\test\\test.txt", []byte{1, 2}, false)
	if err != nil {
		t.Errorf("WriteBytesToFile() error = %v", err)
		return
	}
	t.Log("success")
}

func TestWriteStringToFile(t *testing.T) {
	err := WriteStringToFile("D:\\temp\\test\\test.txt", "hello", false)
	if err != nil {
		t.Errorf("WriteStringToFile() error = %v", err)
		return
	}
	t.Log("success")

	// 测试重复写入是否会丢弃之前的内容
	err = WriteStringToFile("D:\\temp\\test\\test.txt", "ll", false)
	if err != nil {
		t.Errorf("WriteStringToFile() error = %v", err)
		return
	}
	t.Log("success")
}
