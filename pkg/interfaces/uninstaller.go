package interfaces

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// Uninstaller defines the lifecycle steps for uninstalling a service.
type Uninstaller interface {
	PreUninstall(ctx context.Context, svcCtx *types.ServiceContext) error
	Uninstall(ctx context.Context, svcCtx *types.ServiceContext) error
	PostUninstall(ctx context.Context, svcCtx *types.ServiceContext) error
}
