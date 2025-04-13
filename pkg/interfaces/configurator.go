package interfaces

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// Configurator defines the configuration lifecycle for a service.
type Configurator interface {
	PreConfigure(ctx context.Context, svcCtx *types.ServiceContext) error
	Configure(ctx context.Context, svcCtx *types.ServiceContext) error
	PostConfigure(ctx context.Context, svcCtx *types.ServiceContext) error
}
