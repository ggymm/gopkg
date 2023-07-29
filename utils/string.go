package utils

import "strconv"

func ToString(i interface{}) string {
	if i == nil {
		return ""
	}

	switch i.(type) {
	case int:
		return strconv.Itoa(i.(int))
	case int32:
		return strconv.FormatInt(int64(i.(int32)), 10)
	case int64:
		return strconv.FormatInt(i.(int64), 10)
	case float32:
		return strconv.FormatFloat(float64(i.(float32)), 'f', -1, 32)
	case float64:
		return strconv.FormatFloat(i.(float64), 'f', -1, 64)
	case string:
		return i.(string)
	case []byte:
		return string(i.([]byte))
	case bool:
		if i.(bool) {
			return "true"
		} else {
			return "false"
		}
	default:
		return ""
	}
}
