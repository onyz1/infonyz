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

	log.Debug("This is a debug message", &infonyz.Field{
		Key: "user_id",
		Val: 12345,
	})

	log.Info("This is an info message", infonyz.String("operation", "user_login"))

	log.Warn("This is a warning message", &infonyz.Field{
		Key: "disk_space",
		Val: 75.5,
	})

	log.Error("This is an error message", &infonyz.Field{
		Key: "error_code",
		Val: []byte("internal_server_error"),
	})
}
