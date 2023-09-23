package utils

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	t.Log(Now())
}

func TestDuration(t *testing.T) {
	t.Log(time.Duration(30) * time.Minute)
}

func TestYearToSecond(t *testing.T) {
	t.Log(YearToSecond(1))
}

func TestFormatMilli(t *testing.T) {
	t.Log(FormatMilli("2006-01-02 15:04:05", 1694397456620))
}
