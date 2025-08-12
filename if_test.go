package gopkg

import (
	"testing"
)

func Test_If(t *testing.T) {
	t.Log(If(true, 1, 2))
}
