package log

import (
	"log/slog"
	"testing"
)

func Test_Init(t *testing.T) {
	Init("app.log")

	slog.Info("hello")
}
