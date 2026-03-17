package infonyz

// Backend represents the logging backend to use for log messages.
type Backend int

// Logging backends.
const (
	Charm Backend = iota
	Zap
)

// Config holds the configuration options for creating a new logger.
type Config struct {
	// Backend specifies the logging backend to use (e.g., Charm or Zap).
	Backend Backend
	// Level sets the logging level for the logger (e.g., Debug, Info, Warn, Error).
	Level Level
}
