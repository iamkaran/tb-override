// Package logger provides a the method New() which returns a *slog.Logger serving as a common logger throughout the project
package logger

import (
	"context"
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/iamkaran/tb-override/internal/core"
)

func New(level string, format string) *slog.Logger {
	var l slog.Level
	switch level {
	case "debug":
		l = slog.LevelDebug
	case "warn":
		l = slog.LevelWarn
	case "error":
		l = slog.LevelError
	default:
		l = slog.LevelInfo
	}

	replaceAttr := func(groups []string, a slog.Attr) slog.Attr {
		// Check if the attribute key is the source key
		if a.Key == slog.SourceKey {
			// Try to get the source value as a *slog.Source
			if source, ok := a.Value.Any().(*slog.Source); ok {
				// Use filepath.Base to get only the file name
				source.File = filepath.Base(source.File)
				// Remove the function name field if it exists, as requested
				source.Function = ""
			}
		}
		return a
	}

	opts := &slog.HandlerOptions{
		Level:       l,
		AddSource:   true, // Enable source location tracking
		ReplaceAttr: replaceAttr,
	}
	var handler slog.Handler

	if format == "text" {
		handler = slog.NewTextHandler(os.Stdout, opts)
	} else {
		handler = slog.NewJSONHandler(os.Stdout, opts)
	}

	return slog.New(handler)
}

func FromContext(ctx context.Context) *slog.Logger {
	l, ok := ctx.Value(core.LoggerKey).(*slog.Logger)
	if !ok {
		log.Fatalf("Logger could not be fetched: %v\n", ok)
		return slog.Default()
	}

	return l
}
