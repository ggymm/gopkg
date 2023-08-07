package cast

import (
	"encoding/binary"
	"math"
)

func ToBytes(value interface{}) []byte {
	var b []byte
	switch v := value.(type) {
	case string:
		return []byte(v)
	case int:
		b = make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(v))
	case int16:
		b = make([]byte, 2)
		binary.BigEndian.PutUint16(b, uint16(v))
	case int32:
		b = make([]byte, 4)
		binary.BigEndian.PutUint32(b, uint32(v))
	case int64:
		b = make([]byte, 8)
		binary.BigEndian.PutUint64(b, uint64(v))
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
	case float32:
		b = make([]byte, 4)
		binary.BigEndian.PutUint32(b, math.Float32bits(v))
	case float64:
		b = make([]byte, 8)
		binary.BigEndian.PutUint64(b, math.Float64bits(v))
	case bool:
		if v {
			b = []byte{1}
		} else {
			b = []byte{0}
		}
	}
	return b
}
