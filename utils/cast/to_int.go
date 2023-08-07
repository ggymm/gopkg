package cast

import (
	"encoding/binary"
	"fmt"
	"github.com/ggymm/gopkg/types"
	"github.com/pkg/errors"
	"strconv"
)

func toSignedE[T types.Signed](value any, options ...Options) (T, error) {
	var zero T
	var errStr = fmt.Sprintf("unable to cast %#v of type %T to %T", value, value, zero)
	switch v := value.(type) {
	case float32:
		return T(v), nil
	case float64:
		return T(v), nil
	case int:
		return T(v), nil
	case int16:
		return T(v), nil
	case int32:
		return T(v), nil
	case int64:
		return T(v), nil
	case uint:
		return T(v), nil
	case uint16:
		return T(v), nil
	case uint32:
		return T(v), nil
	case uint64:
		return T(v), nil
	case string:
		if v == "" {
			switch options[0].Empty {
			case ValueToZero:
				return zero, nil
			case ValueToError:
				return zero, errors.New(errStr)
			}
		} else {
			// int，int16，int32，int64 转 string
			vv, err := strconv.ParseInt(trimZero(v), 0, 0)
			if err != nil {
				return zero, err
			}
			return T(vv), nil
		}
	case []byte:
		switch len(v) {
		case 2:
			return T(binary.BigEndian.Uint16(v)), nil
		case 4:
			return T(binary.BigEndian.Uint32(v)), nil
		case 8:
			return T(binary.BigEndian.Uint64(v)), nil
		case 0:
			switch options[0].Empty {
			case ValueToZero:
				return zero, nil
			case ValueToError:
				return zero, errors.New(errStr)
			}
		default:
			return zero, errors.New(errStr)
		}
	case bool:
		if v {
			return T(1), nil
		} else {
			return zero, nil
		}
	case nil:
		switch options[0].Nil {
		case NilToZero:
			return zero, nil
		case NilToError:
			return zero, errors.New(errStr)
		}
	}
	return zero, nil
}

// trimZero
// 移除小数部分：123.0 -> 123
func trimZero(s string) string {
	var foundZero bool
	for i := len(s); i > 0; i-- {
		switch s[i-1] {
		case '.':
			if foundZero {
				return s[:i-1]
			}
		case '0':
			foundZero = true
		default:
			return s
		}
	}
	return s
}
