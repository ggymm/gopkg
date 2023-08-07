package utils

import "time"

func Now() int64 {
	return time.Now().UnixMilli()
}

func CurrentTimestamp() int64 {
	return time.Now().UnixMilli()
}
