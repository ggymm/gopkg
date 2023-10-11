package file

import (
	"os"
	"path/filepath"
)

// WriteBytes write bytes to file
func WriteBytes(path string, content []byte, append bool) error {
	var flag int
	if append {
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	} else {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}

	// 判断目录是否存在，不存在则创建
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 打开文件
	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}

	// 写入 字符串 到文件
	_, err = f.Write(content)
	if err != nil {
		return err
	}
	return f.Close()
}

// WriteString write string to file
func WriteString(path string, content string, append bool) error {
	var flag int
	if append {
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	} else {
		flag = os.O_RDWR | os.O_CREATE | os.O_TRUNC
	}

	// 判断目录是否存在，不存在则创建
	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}

	// 打开文件
	f, err := os.OpenFile(path, flag, 0644)
	if err != nil {
		return err
	}

	// 写入 字符串 到文件
	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return f.Close()
}
