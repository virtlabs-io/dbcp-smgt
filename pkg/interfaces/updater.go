package interfaces

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// Updater defines the operations for updating or upgrading a service.
type Updater interface {
	PreUpdate(ctx context.Context, svcCtx *types.ServiceContext) error
	Update(ctx context.Context, svcCtx *types.ServiceContext) error
	PostUpdate(ctx context.Context, svcCtx *types.ServiceContext) error
}
