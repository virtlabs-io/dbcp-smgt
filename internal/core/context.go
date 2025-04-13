package core

import "go.uber.org/zap"

// ServiceContext is used by each lifecycle phase.
type ServiceContext struct {
	Version string
	Config  map[string]interface{}
	Logger  *zap.Logger
	DryRun  bool
}
