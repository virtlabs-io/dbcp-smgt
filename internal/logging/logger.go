// internal/logging/logger.go
package logging

import "go.uber.org/zap"

// NewLogger creates and returns a new zap.Logger instance.
func NewLogger() *zap.Logger {
	logger, _ := zap.NewProduction() // In production code, handle errors properly.
	return logger
}
