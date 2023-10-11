package file

import (
	"testing"
)

func TestWriteBytesToFile(t *testing.T) {
	err := WriteBytes("D:\\temp\\test\\test.txt", []byte{1, 2}, false)
	if err != nil {
		t.Errorf("WriteBytes() error = %v", err)
		return
	}
	t.Log("success")
}

func TestWriteStringToFile(t *testing.T) {
	err := WriteString("D:\\temp\\test\\test.txt", "hello", false)
	if err != nil {
		t.Errorf("WriteString() error = %v", err)
		return
	}
	t.Log("success")

	// 测试重复写入是否会丢弃之前的内容
	err = WriteString("D:\\temp\\test\\test.txt", "ll", false)
	if err != nil {
		t.Errorf("WriteString() error = %v", err)
		return
	}
	t.Log("success")
}
