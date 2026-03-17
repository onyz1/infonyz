package infonyz

import (
	clog "github.com/charmbracelet/log"
	"go.uber.org/zap/zapcore"
)

// Level represents the logging level for log messages.
type Level int

// Logging levels.
const (
	DebugLevel Level = iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

// String returns the string representation of the logging level.
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "debug"
	case InfoLevel:
		return "info"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	default:
		return "info"
	}
}

func (l Level) zapLevel() zapcore.Level {
	switch l {
	case DebugLevel:
		return zapcore.DebugLevel
	case InfoLevel:
		return zapcore.InfoLevel
	case WarnLevel:
		return zapcore.WarnLevel
	case ErrorLevel:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

func (l Level) charmLevel() clog.Level {
	switch l {
	case DebugLevel:
		return clog.DebugLevel
	case InfoLevel:
		return clog.InfoLevel
	case WarnLevel:
		return clog.WarnLevel
	case ErrorLevel:
		return clog.ErrorLevel
	default:
		return clog.InfoLevel
	}
}
