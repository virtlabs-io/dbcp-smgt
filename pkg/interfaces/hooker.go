// pkg/interfaces/hooker.go
package interfaces

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// Hooker is an optional interface for injecting hook behaviors.
type Hooker interface {
	PreInstall(ctx context.Context, svcCtx *types.ServiceContext) error
	PostStart(ctx context.Context, svcCtx *types.ServiceContext) error
	// Define additional hooks as needed.
}
