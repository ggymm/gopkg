package utils

import (
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	t.Log(time.Duration(30) * time.Minute)
}
