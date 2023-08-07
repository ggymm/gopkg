package log

import (
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	testDir, err := os.MkdirTemp("", "test-zerolog-*")
	if err != nil {
		t.Error(err)
		return
	}
	Init(testDir + "/test.log")

	Trace().Msg("trace")
	Debug().Msg("debug")
	Info().Msg("info")
	Warn().Msg("warn")
	Error().Msg("error")
	defer func() {
		Fatal().Msg("fatal")
		Panic().Msg("panic")
	}()
}
