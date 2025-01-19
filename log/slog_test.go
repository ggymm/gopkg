package log

import (
	"log/slog"
	"testing"
)

func Test_Init(t *testing.T) {
	Init("app.log", slog.LevelDebug)

	slog.Info("hello", "world", "111")
	slog.Debug("hello", "world", "222")
	slog.Error("hello", "world", "333")
	slog.Warn("hello", "world", "444")
}
