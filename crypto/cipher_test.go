package crypto

import (
	"testing"
)

func Test_generateAesKey(t *testing.T) {
	t.Log(string(formatDesKey([]byte("ninelockninelock1111111"))))
}

func Test_randomIV(t *testing.T) {
	for i := 0; i < 16; i++ {
		t.Log(randomIV())
	}
}
