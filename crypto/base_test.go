package crypto

import "testing"

func TestBase64(t *testing.T) {
	src := "hello world"
	dst := Base64StdEncode(src)
	t.Log(dst)

	decode := Base64StdDecode(dst)
	if decode != src {
		t.Fatalf("decode is %s, want %s", decode, src)
	}
	t.Log("success")
}

func TestMD5File(t *testing.T) {
	t.Log(MD5File("base_test.go"))
}

func TestMD5String(t *testing.T) {
	src := "hello world"
	dst := MD5String(src)
	t.Log(dst)
}

func TestSha(t *testing.T) {
	src := "hello world"
	t.Log(Sha1String(src))
	t.Log(Sha256String(src))
	t.Log(Sha512String(src))
}
