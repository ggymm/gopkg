package cryptor

import (
	"crypto/cipher"
	"crypto/des"
)

// ------------------------------------------------ default ----------------------------------------------------------------

func DesEncrypt(key, src []byte) []byte {
	return DesEncryptECB(key, src)
}

func DesDecrypt(key, src []byte) []byte {
	return DesDecryptECB(key, src)
}

// ------------------------------------------------ des ecb ----------------------------------------------------------------

// DesEncryptECB encrypts src with key using DES in ECB mode.
// The length of key must be 8 bytes.
func DesEncryptECB(key, src []byte) []byte {
	c, _ := des.NewCipher(formatDesKey(key))
	l := (len(src) + des.BlockSize) / des.BlockSize
	plain := make([]byte, l*des.BlockSize)

	copy(plain, src)

	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}

	dst := make([]byte, len(plain))
	for bs, be := 0, c.BlockSize(); bs <= len(src); bs, be = bs+c.BlockSize(), be+c.BlockSize() {
		c.Encrypt(dst[bs:be], plain[bs:be])
	}
	return dst
}

// DesDecryptECB decrypts src with key using DES in ECB mode.
// The length of key must be 8 bytes.
func DesDecryptECB(key, src []byte) []byte {
	c, _ := des.NewCipher(formatDesKey(key))

	dst := make([]byte, len(src))
	for bs, be := 0, c.BlockSize(); bs < len(src); bs, be = bs+c.BlockSize(), be+c.BlockSize() {
		c.Decrypt(dst[bs:be], src[bs:be])
	}

	trim := 0
	if len(dst) > 0 {
		trim = len(dst) - int(dst[len(dst)-1])
	}
	return dst[:trim]
}

// ------------------------------------------------ des cbc ----------------------------------------------------------------

// DesEncryptCBC encrypts src with key using DES in CBC mode.
// The length of key must be 8 bytes.
// iv is the initialization vector. The length of iv must be 8 bytes.
func DesEncryptCBC(key, src []byte, iv []byte) []byte {
	c, _ := des.NewCipher(formatDesKey(key))
	plain := pkcs7Padding(src, des.BlockSize)

	dst := make([]byte, len(plain))
	cipher.NewCBCEncrypter(c, iv).CryptBlocks(dst, plain)
	return dst
}

// DesDecryptCBC decrypts src with key using DES in CBC mode.
// The length of key must be 8 bytes.
// iv is the initialization vector. The length of iv must be 8 bytes.
func DesDecryptCBC(key, src []byte, iv []byte) []byte {
	c, _ := des.NewCipher(formatDesKey(key))

	iv = formatDesKey(iv)
	dst := make([]byte, len(src))
	cipher.NewCBCDecrypter(c, iv).CryptBlocks(dst, src)
	return pkcs7UnPadding(dst)
}

// ------------------------------------------------ des cfb ----------------------------------------------------------------

// DesEncryptCFB encrypts src with key using DES in CFB mode.
// The length of key must be 8 bytes.
// iv is the initialization vector. The length of iv must be 8 bytes.
func DesEncryptCFB(key, src []byte, iv []byte) []byte {
	c, _ := des.NewCipher(formatDesKey(key))

	dst := make([]byte, len(src))
	cipher.NewCFBEncrypter(c, iv).XORKeyStream(dst, src)
	return dst
}

// DesDecryptCFB decrypts src with key using DES in CFB mode.
// The length of key must be 8 bytes.
// iv is the initialization vector. The length of iv must be 8 bytes.
func DesDecryptCFB(key, src []byte, iv []byte) []byte {
	c, _ := des.NewCipher(formatDesKey(key))

	dst := make([]byte, len(src))
	cipher.NewCFBDecrypter(c, iv).XORKeyStream(dst, src)
	return dst
}
