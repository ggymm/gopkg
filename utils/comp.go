package utils

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"io"
)

func ZLibDecompress(data []byte) ([]byte, error) {
	buf := bytes.NewReader(data)
	reader := flate.NewReader(buf)
	defer func() {
		_ = reader.Close()
	}()
	return io.ReadAll(reader)
}

func GZipDecompress(data []byte) ([]byte, error) {
	buf := bytes.NewBuffer(data)
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return nil, err
	}
	defer func() {
		_ = reader.Close()
	}()
	return io.ReadAll(reader)
}
