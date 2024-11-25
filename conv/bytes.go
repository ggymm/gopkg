package conv

import (
	"encoding/binary"
)

func ParseBytes(value any) []byte {
	var b []byte
	switch v := value.(type) {
	case uint:
		b = make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(v))
	case uint16:
		b = make([]byte, 2)
		binary.BigEndian.PutUint16(b, v)
	case uint32:
		b = make([]byte, 4)
		binary.BigEndian.PutUint32(b, v)
	case uint64:
		b = make([]byte, 8)
		binary.BigEndian.PutUint64(b, v)
	}
	return b
}
