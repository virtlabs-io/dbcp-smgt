package interfaces

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// Validator defines methods for validating different aspects of a service.
type Validator interface {
	ValidateInstallation(ctx context.Context, svcCtx *types.ServiceContext) error
	ValidateConfiguration(ctx context.Context, svcCtx *types.ServiceContext) error
	ValidateRuntime(ctx context.Context, svcCtx *types.ServiceContext) error
}
