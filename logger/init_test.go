package logger

import (
	"testing"
)

func TestInit(t *testing.T) {
	log := Init("test.log")
	log.Info().Msg("info msg")
}
