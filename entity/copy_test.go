package entity

import (
	"testing"
)

type Src struct {
	Name string
	Age  int
}

type Dst struct {
	Name string
	Age  int
}

func Test_Copy(t *testing.T) {
	src := Src{
		Name: "test",
		Age:  18,
	}

	dst := Copy(src, Dst{})
	t.Log(dst)
}
