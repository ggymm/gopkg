package uuid

import (
	"encoding/hex"
)

func NewUUID() string {
	id := New()
	return hex.EncodeToString(id[:])
}
