package utils

import (
	"bufio"
	"io"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func CurrentPath() string {
	var absPath string
	_, filename, _, ok := runtime.Caller(1)
	if ok {
		absPath = path.Dir(filename)
	}

	return absPath
}

// Mkdir create dir if not exists
func Mkdir(dir string) (err error) {
	if _, err = os.Stat(dir); err != nil {
		err = os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}

// ReadFileToLines
// 读取文件内容，按行返回
func ReadFileToLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = f.Close()
	}()

	buf := bufio.NewReader(f)
	lines := make([]string, 0)
	for {
		l, _, err1 := buf.ReadLine()
		if err1 == io.EOF {
			break
		}
		if err1 != nil {
			continue
		}
		lines = append(lines, string(l))
	}
	return lines, nil
}

// ReadFileToString
// 读取文件内容，返回字符串
func ReadFileToString(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// ReadMultipartFileToBytes
// 读取 multipart.FileHeader 文件内容，返回 []byte
func ReadMultipartFileToBytes(fh *multipart.FileHeader) ([]byte, error) {
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

// WriteBytesToFile write bytes to file
func WriteBytesToFile(path string, content []byte, append bool) error {
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

// WriteStringToFile write string to file
func WriteStringToFile(path string, content string, append bool) error {
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
