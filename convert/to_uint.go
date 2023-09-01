package convert

import (
	"encoding/binary"
	"fmt"
	"github.com/ggymm/gopkg/types"
	"strconv"
)

func ToUnsignedE[T types.Unsigned](value any, options ...Options) (T, error) {
	var zero T
	var errStr = fmt.Sprintf("unable to cast %#v of type %T to %T", value, value, zero)
	switch v := value.(type) {
	case float32:
		return signedCheck[float32, T](v)
	case float64:
		return signedCheck[float64, T](v)
	case int:
		return signedCheck[int, T](v)
	case int16:
		return signedCheck[int16, T](v)
	case int32:
		return signedCheck[int32, T](v)
	case int64:
		return signedCheck[int64, T](v)
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
				return zero, fmt.Errorf(errStr)
			}
		} else {
			// uint，uint16，uint32，uint64 转 string
			vv, err := strconv.ParseUint(trimZero(v), 0, 0)
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
				return zero, fmt.Errorf(errStr)
			}
		default:
			return zero, fmt.Errorf(errStr)
		}
	case bool:
		if v {
			return T(1), nil
		} else {
			return zero, nil
		}
	case nil:
		switch options[0].Empty {
		case ValueToZero:
			return zero, nil
		case ValueToError:
			return zero, fmt.Errorf(errStr)
		}
	}
	return zero, nil
}

func signedCheck[Source types.Signed | types.Float, Target types.Unsigned](source Source) (Target, error) {
	if source < 0 {
		return 0, fmt.Errorf("unable to cast negative value")
	} else {
		return Target(source), nil
	}
}
