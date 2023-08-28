package cryptor

import (
	"crypto/aes"
	"crypto/cipher"
)

// ------------------------------------------------ default ----------------------------------------------------------------

func AesEncrypt(key, src []byte) []byte {
	dst, _ := AesEncryptE(key, src)
	return dst
}

func AesEncryptE(key, src []byte) ([]byte, error) {
	return AesEncryptECB(key, src)
}

func AesDecrypt(key, src []byte) []byte {
	dst, _ := AesDecryptE(key, src)
	return dst
}

func AesDecryptE(key, src []byte) ([]byte, error) {
	return AesDecryptECB(key, src)
}

// ------------------------------------------------ aes ecb ----------------------------------------------------------------

// AesEncryptECB encrypts src with key using AES in ECB mode.
// The length of key must be 16, 24 or 32 bytes to select
// AES-128, AES-192 or AES-256.
func AesEncryptECB(key, src []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	l := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, l*aes.BlockSize)

	copy(plain, src)

	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}

	dst := make([]byte, len(plain))
	for bs, be := 0, c.BlockSize(); bs <= len(src); bs, be = bs+c.BlockSize(), be+c.BlockSize() {
		c.Encrypt(dst[bs:be], plain[bs:be])
	}
	return dst, nil
}

// AesDecryptECB decrypts src with key using AES in ECB mode.
// The length of key must be 16, 24 or 32 bytes to select
// AES-128, AES-192 or AES-256.
func AesDecryptECB(key, src []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(src))
	for bs, be := 0, c.BlockSize(); bs < len(src); bs, be = bs+c.BlockSize(), be+c.BlockSize() {
		c.Decrypt(dst[bs:be], src[bs:be])
	}

	trim := 0
	if len(dst) > 0 {
		trim = len(dst) - int(dst[len(dst)-1])
	}
	return dst[:trim], nil
}

// ------------------------------------------------ aes cbc ----------------------------------------------------------------

// AesEncryptCBC encrypts src with key using AES in CBC mode.
// The length of key must be 16, 24 or 32 bytes to select
// AES-128, AES-192 or AES-256.
// iv is the key to encrypt the first block.
func AesEncryptCBC(key, src []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	plain := pkcs7Padding(src, aes.BlockSize)

	dst := make([]byte, len(plain))
	cipher.NewCBCEncrypter(c, key[:aes.BlockSize]).CryptBlocks(dst, plain)
	return dst, nil
}

// AesDecryptCBC decrypts src with key using AES in CBC mode.
// The length of key must be 16, 24 or 32 bytes to select
// AES-128, AES-192 or AES-256.
// iv is the key to decrypt the first block.
func AesDecryptCBC(key, src []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(src))
	cipher.NewCBCDecrypter(c, key[:aes.BlockSize]).CryptBlocks(dst, src)
	return pkcs7UnPadding(dst), nil
}

// ------------------------------------------------ aes cfb ----------------------------------------------------------------

// AesEncryptCFB encrypts src with key using AES in CFB mode.
// The length of key must be 16, 24 or 32 bytes to select
// AES-128, AES-192 or AES-256.
// iv is the key to encrypt the first block.
func AesEncryptCFB(key, src []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(src))
	cipher.NewCFBEncrypter(c, key[:aes.BlockSize]).XORKeyStream(dst, src)
	return dst, nil
}

// AesDecryptCFB decrypts src with key using AES in CFB mode.
// The length of key must be 16, 24 or 32 bytes to select
// AES-128, AES-192 or AES-256.
// iv is the key to decrypt the first block.
func AesDecryptCFB(key, src []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dst := make([]byte, len(src))
	cipher.NewCFBDecrypter(c, key[:aes.BlockSize]).XORKeyStream(dst, src)
	return dst, nil
}
