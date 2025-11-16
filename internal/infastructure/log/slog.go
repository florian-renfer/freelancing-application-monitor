package log

import (
	"log/slog"
	"os"

	"github.com/florian-renfer/freelancing-application-monitor/internal/adapter/logger"
)

type slogLogger struct {
	logger *slog.Logger
}

func NewSlogLogger() (logger.Logger, error) {
	handler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	log := slog.New(handler)
	return &slogLogger{logger: log}, nil
}

func (l *slogLogger) Infof(format string, args ...any) {
	l.logger.Info(format, args...)
}

func (l *slogLogger) Warnf(format string, args ...any) {
	l.logger.Warn(format, args...)
}

func (l *slogLogger) Errorf(format string, args ...any) {
	l.logger.Error(format, args...)
}

func (l *slogLogger) WithFields(fields logger.Fields) logger.Logger {
	var f = make([]interface{}, 0)
	for index, field := range fields {
		f = append(f, index)
		f = append(f, field)
	}

	log := l.logger.With(f...)
	return &slogLogger{logger: log}
}

func (l *slogLogger) WithError(err error) logger.Logger {
	var log = l.logger.With(err.Error())
	return &slogLogger{logger: log}
}
