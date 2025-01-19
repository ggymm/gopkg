package log

import (
	"io"
	"log/slog"
	"os"
	"time"

	"github.com/ggymm/gopkg/rolling"
)

func New(filename string, level slog.Level) *slog.Logger {
	writer := io.MultiWriter(
		&rolling.Logger{
			Filename:   filename,
			MaxAge:     30,  // days
			MaxSize:    256, // megabytes
			MaxBackups: 128, // files
		},
		io.MultiWriter(os.Stdout),
	)
	opt := &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		ReplaceAttr: func(groups []string, attr slog.Attr) slog.Attr {
			if attr.Key == slog.TimeKey {
				if t, ok := attr.Value.Any().(time.Time); ok {
					return slog.Attr{
						Key:   attr.Key,
						Value: slog.StringValue(t.Format("2006-01-02 15:04:05")),
					}
				}
			}
			return attr
		},
	}
	return slog.New(slog.NewTextHandler(writer, opt))
}

func Init(filename string, level slog.Level) {
	slog.SetDefault(New(filename, level))
}
