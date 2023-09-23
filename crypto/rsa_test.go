package crypto

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"testing"
)

func TestRsa(t *testing.T) {
	// 生成私钥
	priKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}
	privateKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(priKey),
	})

	// 生成公钥
	pubKey, err := x509.MarshalPKIXPublicKey(&priKey.PublicKey)
	if err != nil {
		t.Fatal(err)
	}
	publicKey := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubKey,
	})

	// 加密
	src := []byte("hello world")
	encrypted, err := RsaEncrypt(src, publicKey)
	if err != nil {
		t.Fatal(err)
	}

	// 解密
	decrypted, err := RsaDecrypt(encrypted, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	if string(decrypted) != string(src) {
		t.Fatal("decrypted != src")
	}
	t.Log("success")
}
