// pkg/types/context.go
package types

import "go.uber.org/zap"

// ServiceContext holds configuration shared across lifecycle phases.
type ServiceContext struct {
	Version string
	Config  map[string]interface{}
	Logger  *zap.Logger
	DryRun  bool
}
