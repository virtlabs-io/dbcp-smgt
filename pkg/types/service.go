// pkg/types/service.go
package types

import "context"

// (if needed; otherwise leave it as-is)

// Service defines the base service interface.
type Service interface {
	Name() string
	Version() string
	RequiredInterfaces() []string // Optional: list implemented interfaces.
}

// BaseService provides default (no-op) hook implementations.
type BaseService struct{}

func (b *BaseService) PreInstall(ctx context.Context, svcCtx *ServiceContext) error {
	// Default no-op implementation.
	return nil
}
