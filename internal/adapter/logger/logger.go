package logger

type Logger interface {
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	WithFields(keyValues Fields) Logger
	WithError(err error) Logger
}

type Fields map[string]any
