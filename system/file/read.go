package file

import (
	"bufio"
	"io"
	"os"
)

// ReadLines
// 读取文件内容，按行返回
func ReadLines(path string) ([]string, error) {
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

// ReadString
// 读取文件内容，返回字符串
func ReadString(path string) (string, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(file), nil
}

// ReadBytes
// 读取文件内容，返回 []byte
func ReadBytes(path string) ([]byte, error) {
	return os.ReadFile(path)
}
