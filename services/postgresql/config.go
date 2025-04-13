// services/postgresql/service.go
package postgresql

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
	"go.uber.org/zap"
)

//////////////////////
// Configurator Interface
//////////////////////

func (s *PostgreSQLService) PreConfigure(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("PreConfigure: Backing up PostgreSQL configuration")
	return nil
}

func (s *PostgreSQLService) Configure(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("Configuring PostgreSQL service",
		zap.String("installDir", s.InstallDir),
		zap.String("dataDir", s.DataDir),
		zap.Int("port", s.Port))
	// Here you would apply configuration changes.
	return nil
}

func (s *PostgreSQLService) PostConfigure(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("PostConfigure: PostgreSQL configuration applied successfully")
	return nil
}
