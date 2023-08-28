package utils

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	t.Log(time.Duration(30) * time.Minute)
}

func TestYearToSecond(t *testing.T) {
	t.Log(YearToSecond(1))
}
