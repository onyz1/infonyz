package infonyz

import (
	"io"
)

// Logger defines the interface for logging messages at various levels with optional fields.
type Logger interface {
	// Debug logs a message at the Debug level with optional fields.
	Debug(msg string, fields ...*Field)
	// Info logs a message at the Info level with optional fields.
	Info(msg string, fields ...*Field)
	// Warn logs a message at the Warn level with optional fields.
	Warn(msg string, fields ...*Field)
	// Error logs a message at the Error level with optional fields.
	Error(msg string, fields ...*Field)

	// SetLevel sets the logging level for the logger.
	SetLevel(Level)
	// GetLevel returns the current logging level of the logger.
	GetLevel() Level
	// IsLevel checks if the provided logging level is enabled for the logger.
	IsLevel(Level) bool
}

type noopLogger struct{}

func (n *noopLogger) Debug(msg string, fields ...*Field) {}
func (n *noopLogger) Info(msg string, fields ...*Field)  {}
func (n *noopLogger) Warn(msg string, fields ...*Field)  {}
func (n *noopLogger) Error(msg string, fields ...*Field) {}

func (n *noopLogger) SetLevel(l Level)     {}
func (n *noopLogger) GetLevel() Level      { return ErrorLevel }
func (n *noopLogger) IsLevel(l Level) bool { return false }

// NoopLogger returns a logger that discards all log messages. This is useful as a default logger when no other logger is configured.
func NoopLogger() Logger {
	return &noopLogger{}
}

// New creates a new logger based on the provided configuration and output writer. If the writer is nil, it defaults to io.Discard.
func New(cfg *Config, w io.Writer) Logger {
	if w == nil {
		w = io.Discard
	}

	switch cfg.Backend {
	case Charm:
		return newCharm(cfg, w)
	case Zap:
		return newZap(cfg, w)
	default:
		return newCharm(cfg, w)
	}
}
