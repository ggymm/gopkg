package files

import (
	"io"
	"mime/multipart"
)

// ReadMultipartFileBytes
// 读取 multipart.FileHeader 文件内容，返回 []byte
func ReadMultipartFileBytes(fh *multipart.FileHeader) ([]byte, error) {
	var (
		err     error
		file    multipart.File
		content []byte
	)

	// 打开文件
	file, err = fh.Open()
	if err != nil {
		return nil, err
	}
	content, err = io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	defer func() {
		e := file.Close()
		if err == nil {
			err = e
		}
	}()

	return content, err
}
