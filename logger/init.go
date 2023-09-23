package logger

import (
	"io"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)

func Init(filename ...string) zerolog.Logger {
	writers := io.MultiWriter(console())
	if len(filename) > 0 {
		if len(filename[0]) > 0 {
			writers = io.MultiWriter(writers, rollingFile(filename[0]))
		}
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
