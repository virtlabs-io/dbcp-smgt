// pkg/interfaces/installer.go
package interfaces

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// Installer defines the installation lifecycle phase.
type Installer interface {
	Install(ctx context.Context, svcCtx *types.ServiceContext) error
	DryRunInstall(ctx context.Context, svcCtx *types.ServiceContext) ([]string, error)
	PreInstall(ctx context.Context, svcCtx *types.ServiceContext) error
	PostInstall(ctx context.Context, svcCtx *types.ServiceContext) error
}
