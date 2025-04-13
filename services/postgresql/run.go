// services/postgresql/service.go
package postgresql

import (
	"context"
	"time"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
	"go.uber.org/zap"
)

//////////////////////
// Runner Interface
//////////////////////

func (s *PostgreSQLService) PreStart(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("PreStart: Running health checks before starting PostgreSQL")
	return nil
}

func (s *PostgreSQLService) Start(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("Starting PostgreSQL service", zap.String("dataDir", s.DataDir), zap.Int("port", s.Port))
	time.Sleep(2 * time.Second) // Simulate startup delay
	return nil
}

func (s *PostgreSQLService) PostStart(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("PostStart: PostgreSQL service is now running")
	return nil
}

func (s *PostgreSQLService) Stop(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("Stopping PostgreSQL service")
	time.Sleep(1 * time.Second)
	return nil
}

func (s *PostgreSQLService) PostStop(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("PostStop: PostgreSQL service has been stopped")
	return nil
}
