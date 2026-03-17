# infonyz

`infonyz` is a small logger abstraction for Go with two interchangeable backends:

- `Charm` (human-readable terminal logs via `github.com/charmbracelet/log`)
- `Zap` (structured JSON logs via `go.uber.org/zap`)

It provides a single `Logger` interface, a lightweight `Field` type, runtime level control, and context helpers.

## Features

- Unified logger interface across backends
- Structured fields (`String`, `Int`, `Bool`, `Bytes`, etc.)
- Runtime log level updates (`SetLevel`)
- Context integration (`WithLogger`, `FromContext`)
- No-op fallback logger (`NoopLogger`)

## Install

```bash
go get github.com/onyz1/infonyz
```

## Quick Start

```go
package main

import (
	"os"

	"github.com/onyz1/infonyz"
)

func main() {
	log := infonyz.New(&infonyz.Config{
		Backend: infonyz.Zap,
		Level:   infonyz.DebugLevel,
	}, os.Stdout)

	log.Debug("debug message", infonyz.Int("user_id", 12345))
	log.Info("user login", infonyz.String("operation", "login"))
	log.Warn("disk usage high", infonyz.Float64("disk_pct", 75.5))
	log.Error("request failed", infonyz.Bytes("error_code", []byte("internal_server_error")))
}
```

## Backends

### `Charm`

- Readable, styled terminal output
- Good for local development and CLI tools

```go
log := infonyz.New(&infonyz.Config{
	Backend: infonyz.Charm,
	Level:   infonyz.InfoLevel,
}, os.Stdout)
```

### `Zap`

- JSON output suitable for ingestion by log processors
- Good for production services

```go
log := infonyz.New(&infonyz.Config{
	Backend: infonyz.Zap,
	Level:   infonyz.InfoLevel,
}, os.Stdout)
```

## Log Levels

Available levels:

- `infonyz.DebugLevel`
- `infonyz.InfoLevel`
- `infonyz.WarnLevel`
- `infonyz.ErrorLevel`

You can change the level at runtime:

```go
log.SetLevel(infonyz.WarnLevel)

if log.IsLevel(infonyz.DebugLevel) {
	log.Debug("expensive debug message")
}
```

## Fields

Create fields using helpers:

- `infonyz.String(key, value)`
- `infonyz.Int(key, value)`
- `infonyz.Int64(key, value)`
- `infonyz.Float32(key, value)`
- `infonyz.Float64(key, value)`
- `infonyz.Bool(key, value)`
- `infonyz.Bytes(key, value)`
- `infonyz.F(key, value)` for arbitrary values

Example:

```go
log.Info("order processed",
	infonyz.String("order_id", "ORD-123"),
	infonyz.Float64("amount", 49.99),
	infonyz.Bool("paid", true),
)
```

## Context Integration

Attach a logger to `context.Context`:

```go
ctx := infonyz.WithLogger(context.Background(), log)

reqLog := infonyz.FromContext(ctx)
reqLog.Info("processing request")
```

If no logger exists in context, `FromContext` returns a no-op logger.

## API Overview

### Types

- `type Logger interface`
- `type Config struct { Backend Backend; Level Level }`
- `type Field struct { Key string; Val any }`

### Constructors and helpers

- `func New(cfg *Config, w io.Writer) Logger`
- `func NoopLogger() Logger`
- `func WithLogger(ctx context.Context, log Logger) context.Context`
- `func FromContext(ctx context.Context) Logger`

### Logger methods

- `Debug(msg string, fields ...*Field)`
- `Info(msg string, fields ...*Field)`
- `Warn(msg string, fields ...*Field)`
- `Error(msg string, fields ...*Field)`
- `SetLevel(Level)`
- `GetLevel() Level`
- `IsLevel(Level) bool`

## Notes

- If `w` passed to `New` is `nil`, output is discarded (`io.Discard`).
- Unsupported field value types are handled through backend-specific generic encoding.

## License

No license use as you wish.