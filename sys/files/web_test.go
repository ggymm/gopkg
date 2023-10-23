package files

import (
	"mime/multipart"
	"testing"
)

func TestReadMultipartFileToBytes(t *testing.T) {
	// 创建 multipart.FileHeader
	fh := &multipart.FileHeader{
		Filename: "web_test.go",
		// TODO: 构建 multipart.FileHeader
	}
	content, err := ReadMultipartFileBytes(fh)
	if err != nil {
		t.Errorf("ReadMultipartFileBytes() error = %v", err)
		return
	}
	t.Log(content)
}
