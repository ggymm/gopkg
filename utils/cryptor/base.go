package cryptor

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"os"

	"github.com/pkg/errors"
)

func MD5File(path string) (string, error) {
	if f, err := os.Stat(path); err != nil {
		return "", err
	} else if f.IsDir() {
		return "", errors.New("path is a directory")
	}

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		_ = file.Close()
	}()

	h := md5.New()
	chunk := make([]byte, 1024*1024)
	for {
		n, readErr := file.Read(chunk)
		if readErr != nil {
			break
		}
		h.Write(chunk[:n])
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func MD5Bytes(data []byte) string {
	h := md5.New()
	h.Write(data)
	return hex.EncodeToString(h.Sum(nil))
}

func MD5String(src string) string {
	return MD5Bytes([]byte(src))
}

func Sha1Bytes(data []byte) string {
	s := sha1.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}

func Sha1String(src string) string {
	return Sha1Bytes([]byte(src))
}

func Sha256Bytes(data []byte) string {
	s := sha256.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}

func Sha256String(src string) string {
	return Sha256Bytes([]byte(src))
}

func Sha512Bytes(data []byte) string {
	s := sha512.New()
	s.Write(data)
	return hex.EncodeToString(s.Sum(nil))
}

func Sha512String(src string) string {
	return Sha512Bytes([]byte(src))
}

func Base64StdEncode(src string) string {
	return base64.StdEncoding.EncodeToString([]byte(src))
}

func Base64StdDecode(src string) string {
	b, _ := base64.StdEncoding.DecodeString(src)
	return string(b)
}
