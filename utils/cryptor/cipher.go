package cryptor

import (
	"bytes"
	"crypto/rand"
)

func formatDesKey(key []byte) []byte {
	genKey := make([]byte, 8)
	copy(genKey, key)
	for i := 8; i < len(key); {
		for j := 0; j < 8 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

func randomIV() []byte {
	randIV := make([]byte, 8)
	_, _ = rand.Read(randIV)
	return randIV
}

func pkcs7Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padText...)
}

func pkcs7UnPadding(src []byte) []byte {
	length := len(src)
	unPadding := int(src[length-1])
	return src[:(length - unPadding)]
}
