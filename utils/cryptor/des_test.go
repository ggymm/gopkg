package cryptor

import (
	"bytes"
	"testing"
)

func TestDes(t *testing.T) {
	key := []byte("12345678")
	src := []byte("12345678")

	dst := DesEncrypt(key, src)
	t.Log(dst)

	dst = DesDecrypt(key, dst)
	t.Log(dst)

	if bytes.Equal(src, dst) == false {
		t.Fatal("not equal")
	} else {
		t.Log("equal")
	}
}

func TestDesECB(t *testing.T) {
	key := []byte("12345678")
	src := []byte("12345678")

	dst := DesEncryptECB(key, src)
	t.Log(dst)

	dst = DesDecryptECB(key, dst)
	t.Log(dst)

	if bytes.Equal(src, dst) == false {
		t.Fatal("not equal")
	} else {
		t.Log("equal")
	}
}

func TestDesCBC(t *testing.T) {
	iv := []byte("12345678")
	key := []byte("12345678")
	src := []byte("12345678")

	dst := DesEncryptCBC(key, src, iv)
	t.Log(dst)

	dst = DesDecryptCBC(key, dst, iv)
	t.Log(dst)

	if bytes.Equal(src, dst) == false {
		t.Fatal("not equal")
	} else {
		t.Log("equal")
	}
}

func TestDesCFB(t *testing.T) {
	iv := []byte("12345678")
	key := []byte("12345678")
	src := []byte("12345678")

	dst := DesEncryptCFB(key, src, iv)
	t.Log(dst)

	dst = DesDecryptCFB(key, dst, iv)
	t.Log(dst)

	if bytes.Equal(src, dst) == false {
		t.Fatal("not equal")
	} else {
		t.Log("equal")
	}
}
