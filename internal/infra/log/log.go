package log

import (
	"log/slog"
	"os"

	"github.com/felipeversiane/donate-api/internal/infra/config"
)

type Logger struct {
	config config.LogConfig
}

type LoggerInterface interface {
	Configure()
}

func NewLogger(config config.LogConfig) LoggerInterface {
	return &Logger{config}
}

func (l *Logger) Configure() {
	level := getLogLevel(l.config)
	logConfig := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: level}))
	slog.SetDefault(logConfig)
}

func getLogLevel(config config.LogConfig) slog.Level {
	level := config.Level

	switch level {
	case "DEBUG":
		return slog.LevelDebug
	case "INFO":
		return slog.LevelInfo
	case "WARN":
		return slog.LevelWarn
	case "ERROR":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
