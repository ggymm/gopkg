package log

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger zerolog.Logger

func Init(filename string) {
	writers := io.MultiWriter(
		newRollingFile(filename),
		zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.DateTime},
	)

	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger = zerolog.New(writers).With().Caller().Timestamp().Logger()
}

func newRollingFile(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    128, // megabytes
		MaxAge:     30,  // days
		MaxBackups: 100, // files
	}
}
