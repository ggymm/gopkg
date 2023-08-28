package cryptor

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
)

func RsaEncrypt(src, pubKey []byte) ([]byte, error) {
	block, _ := pem.Decode(pubKey)

	key, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.EncryptPKCS1v15(rand.Reader, key.(*rsa.PublicKey), src)
}

func RsaDecrypt(src, priKey []byte) ([]byte, error) {
	block, _ := pem.Decode(priKey)

	key, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return rsa.DecryptPKCS1v15(rand.Reader, key, src)
}
