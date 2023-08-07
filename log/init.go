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
	writers := io.MultiWriter(console())
	if len(filename) > 0 {
		writers = io.MultiWriter(writers, rollingFile(filename))
	}

	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	logger = zerolog.New(writers).With().Caller().Timestamp().Logger()
}

func InitCustom(filename string) zerolog.Logger {
	writers := io.MultiWriter(console())
	if len(filename) > 0 {
		writers = io.MultiWriter(writers, rollingFile(filename))
	}

	zerolog.TimeFieldFormat = time.DateTime
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return zerolog.New(writers).With().Caller().Timestamp().Logger()
}

func console() io.Writer {
	return zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: time.DateTime,
	}
}

func rollingFile(filename string) io.Writer {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    128, // megabytes
		MaxAge:     30,  // days
		MaxBackups: 100, // files
	}
}
