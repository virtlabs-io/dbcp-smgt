package interfaces

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// Runner defines the lifecycle operations for starting and stopping a service.
type Runner interface {
	PreStart(ctx context.Context, svcCtx *types.ServiceContext) error
	Start(ctx context.Context, svcCtx *types.ServiceContext) error
	PostStart(ctx context.Context, svcCtx *types.ServiceContext) error
	Stop(ctx context.Context, svcCtx *types.ServiceContext) error
	PostStop(ctx context.Context, svcCtx *types.ServiceContext) error
}
