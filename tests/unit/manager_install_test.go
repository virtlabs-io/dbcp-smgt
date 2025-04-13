package unit

import (
	"context"
	"testing"

	"go.uber.org/zap/zaptest"

	"github.com/virtlabs-io/dbcp-smgt/internal/core"
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
	"github.com/virtlabs-io/dbcp-smgt/services/postgresql"
)

func TestExecuteInstall_DryRun(t *testing.T) {
	logger := zaptest.NewLogger(t)
	svcCtx := &types.ServiceContext{
		Version: "13.3",
		Config:  map[string]interface{}{"env": "test"},
		Logger:  logger,
		DryRun:  true,
	}

	manager := core.NewServiceManager(svcCtx)
	pgService := &postgresql.PostgreSQLService{
		VersionField:       "13.3",
		InstallDir:         "/opt/postgresql",
		DataDir:            "/var/lib/postgresql/13/main",
		Port:               5432,
		MaxConnections:     200,
		Username:           "postgres",
		Password:           "secret",
		ClusterName:        "test_cluster",
		BackupEnabled:      true,
		BackupDir:          "/backup/postgresql",
		RetentionDays:      7,
		MaintenanceWindow:  "Sunday 02:00",
		ReplicationEnabled: false,
		CacheSize:          "512MB",
		LogDirectory:       "/var/log/postgresql",
		Timezone:           "UTC",
		SSLEnabled:         true,
		Locale:             "en_US.UTF-8",
		Autovacuum:         true,
		DataEncryption:     false,
	}
	if err := manager.ExecuteInstall(context.Background(), pgService); err != nil {
		t.Fatalf("Expected no error in dry-run, but got: %v", err)
	}
}
