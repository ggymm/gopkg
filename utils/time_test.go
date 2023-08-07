package utils

import (
	"testing"
	"time"
)

func TestCurrentTimestamp(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{"test1", time.Now().UnixMilli()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CurrentTimestamp(); got != tt.want {
				t.Errorf("CurrentTimestamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
