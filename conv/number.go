package conv

import (
	"strconv"
)

func ParseInt(s string) int {
	i, err := strconv.ParseInt(s, 10, 32)
	if err != nil {
		return 0
	}
	return int(i)
}

func ParseInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return 0
	}
	return i
}
