package conv

import (
	"testing"
)

func Test_ToInt(t *testing.T) {
	s := "123"
	t.Log(ToInt(s))
	t.Log(ToIntOrDefault(s, 123))
	t.Log(ToInt16(s))
	t.Log(ToInt16OrDefault(s, 123))
	t.Log(ToInt32(s))
	t.Log(ToInt32OrDefault(s, 123))
	t.Log(ToInt64(s))
	t.Log(ToInt64OrDefault(s, 123))
}

func Test_ToUint(t *testing.T) {
	s := "123"
	t.Log(ToUint(s))
	t.Log(ToUintOrDefault(s, 123))
	t.Log(ToUint16(s))
	t.Log(ToUint16OrDefault(s, 123))
	t.Log(ToUint32(s))
	t.Log(ToUint32OrDefault(s, 123))
	t.Log(ToUint64(s))
	t.Log(ToUint64OrDefault(s, 123))
}
