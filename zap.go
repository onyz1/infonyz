package infonyz

import (
	"io"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLogger struct {
	log   *zap.Logger
	level zap.AtomicLevel
}

func newZap(cfg *Config, w io.Writer) Logger {
	level := zap.NewAtomicLevelAt(cfg.Level.zapLevel())

	enc := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	ws := zapcore.AddSync(w)

	core := zapcore.NewCore(enc, ws, level)

	l := zap.New(core)

	return &zapLogger{
		log:   l,
		level: level,
	}
}

func fieldsToZap(fs []*Field) []zap.Field {
	out := make([]zap.Field, 0, len(fs))

	for _, field := range fs {
		switch v := field.Val.(type) {
		case string:
			out = append(out, zap.String(field.Key, v))
		case int64:
			out = append(out, zap.Int64(field.Key, v))
		case int:
			out = append(out, zap.Int(field.Key, v))
		case float64:
			out = append(out, zap.Float64(field.Key, v))
		case float32:
			out = append(out, zap.Float32(field.Key, v))
		case bool:
			out = append(out, zap.Bool(field.Key, v))
		case []byte:
			out = append(out, zap.ByteString(field.Key, v))
		default:
			out = append(out, zap.Any(field.Key, v))
		}
	}
	return out
}

// Debug logs a message at the Debug level with optional fields.
func (z *zapLogger) Debug(msg string, fs ...*Field) {
	z.log.Debug(msg, fieldsToZap(fs)...)
}

// Info logs a message at the Info level with optional fields.
func (z *zapLogger) Info(msg string, fs ...*Field) {
	z.log.Info(msg, fieldsToZap(fs)...)
}

// Warn logs a message at the Warn level with optional fields.
func (z *zapLogger) Warn(msg string, fs ...*Field) {
	z.log.Warn(msg, fieldsToZap(fs)...)
}

// Error logs a message at the Error level with optional fields.
func (z *zapLogger) Error(msg string, fs ...*Field) {
	z.log.Error(msg, fieldsToZap(fs)...)
}

// SetLevel sets the logging level for the logger.
func (z *zapLogger) SetLevel(l Level) {
	z.level.SetLevel(l.zapLevel())
}

// GetLevel returns the current logging level of the logger.
func (z *zapLogger) GetLevel() Level {
	switch z.level.Level() {
	case zap.DebugLevel:
		return DebugLevel
	case zap.InfoLevel:
		return InfoLevel
	case zap.WarnLevel:
		return WarnLevel
	case zap.ErrorLevel:
		return ErrorLevel
	default:
		return InfoLevel
	}
}

// IsLevel checks if the provided logging level is enabled for the logger.
func (z *zapLogger) IsLevel(l Level) bool { return l >= z.GetLevel() }
