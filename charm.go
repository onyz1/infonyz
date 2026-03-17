package infonyz

import (
	"io"

	clog "github.com/charmbracelet/log"
)

type charmLogger struct {
	log   *clog.Logger
	level clog.Level
}

func newCharm(cfg *Config, w io.Writer) Logger {
	level := cfg.Level.charmLevel()

	l := clog.NewWithOptions(w, clog.Options{
		Level:           level,
		ReportTimestamp: true,
	})

	c := &charmLogger{
		log:   l,
		level: level,
	}

	return c
}

func fieldsToCharm(fs []*Field) []any {
	kv := make([]any, 0, len(fs)*2)
	for _, field := range fs {
		kv = append(kv, field.Key, field.Val)
	}
	return kv
}

// Info logs a message at the Info level with optional fields.
func (c *charmLogger) Info(msg string, fs ...*Field) {
	c.log.Info(msg, fieldsToCharm(fs)...)
}

// Debug logs a message at the Debug level with optional fields.
func (c *charmLogger) Debug(msg string, fs ...*Field) {
	c.log.Debug(msg, fieldsToCharm(fs)...)
}

// Warn logs a message at the Warn level with optional fields.
func (c *charmLogger) Warn(msg string, fs ...*Field) {
	c.log.Warn(msg, fieldsToCharm(fs)...)
}

// Error logs a message at the Error level with optional fields.
func (c *charmLogger) Error(msg string, fs ...*Field) {
	c.log.Error(msg, fieldsToCharm(fs)...)
}

// SetLevel sets the logging level for the logger.
func (c *charmLogger) SetLevel(l Level) {
	c.level = l.charmLevel()
	c.log.SetLevel(c.level)
}

// GetLevel returns the current logging level of the logger.
func (c *charmLogger) GetLevel() Level {
	switch c.level {
	case clog.DebugLevel:
		return DebugLevel
	case clog.InfoLevel:
		return InfoLevel
	case clog.WarnLevel:
		return WarnLevel
	case clog.ErrorLevel:
		return ErrorLevel
	default:
		return InfoLevel
	}
}

// IsLevel checks if the provided logging level is enabled for the logger.
func (c *charmLogger) IsLevel(l Level) bool { return l >= c.GetLevel() }
