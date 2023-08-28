package cryptor

import (
	"bytes"
	"encoding/hex"
	"strings"
	"testing"
)

func TestAes(t *testing.T) {
	var (
		key = []byte("ninelockninelock")
		src = []byte("hello world hello world hello world hello world")

		encrypted []byte
		decrypted []byte
	)

	encrypted = AesEncrypt(key, src)
	t.Log(encrypted)
	t.Log(strings.ToUpper(hex.EncodeToString(encrypted)))

	decrypted = AesDecrypt(key, encrypted)
	t.Log(string(decrypted))

	if bytes.Equal(src, decrypted) {
		t.Log("equal")
	} else {
		t.Log("not equal")
	}

	encrypted, _ = hex.DecodeString("14DBA52A343BBD5E4F933D9581B463EF")
	decrypted = AesDecrypt(key, encrypted)
	t.Log(string(decrypted))
}

func TestAesECB(t *testing.T) {
	var (
		key = []byte("ninelockninelock")
		src = []byte("hello world hello world hello world hello world")

		err       error
		encrypted []byte
		decrypted []byte
	)

	encrypted, err = AesEncryptECB(key, src)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encrypted)
	t.Log(strings.ToUpper(hex.EncodeToString(encrypted)))

	decrypted, err = AesDecryptECB(key, encrypted)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(decrypted))

	encrypted, _ = hex.DecodeString("14DBA52A343BBD5E4F933D9581B463EF")
	decrypted, err = AesDecryptECB(key, encrypted)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(decrypted))
}

func TestAesCBC(t *testing.T) {
	var (
		key = []byte("1234567812345678")
		src = []byte("hello world")

		err       error
		encrypted []byte
		decrypted []byte
	)

	encrypted, err = AesEncryptCBC(key, src)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encrypted)
	t.Log(strings.ToUpper(hex.EncodeToString(encrypted)))

	decrypted, err = AesDecryptCBC(key, encrypted)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(decrypted))

	if bytes.Equal(src, decrypted) {
		t.Log("equal")
	} else {
		t.Log("not equal")
	}
}

func TestAesCFB(t *testing.T) {
	var (
		key = []byte("1234567812345678")
		src = []byte("hello world")

		err       error
		encrypted []byte
		decrypted []byte
	)

	encrypted, err = AesEncryptCFB(key, src)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(encrypted)
	t.Log(strings.ToUpper(hex.EncodeToString(encrypted)))

	decrypted, err = AesDecryptCFB(key, encrypted)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(decrypted))

	if bytes.Equal(src, decrypted) {
		t.Log("equal")
	} else {
		t.Log("not equal")
	}
}
