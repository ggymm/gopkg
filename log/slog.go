package log

import (
	"io"
	"log/slog"
	"os"

	"github.com/ggymm/gopkg/rolling"
)

func Init(filename string) {
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
		AddSource: true,
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(writer, opt)))
}
