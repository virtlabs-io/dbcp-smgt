// services/postgresql/service.go
package postgresql

import (
	"context"
	"strings"
	"time"

	"github.com/blang/semver"
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
	"go.uber.org/zap"
)

// PostgreSQLService implements lifecycle interfaces as well as PackageInfoProvider.
type PostgreSQLService struct {
	types.BaseService
	VersionField string

	// Service-specific configuration from the config file.
	InstallDir         string
	DataDir            string
	Port               int
	MaxConnections     int
	Username           string
	Password           string
	ClusterName        string
	BackupEnabled      bool
	BackupDir          string
	RetentionDays      int
	MaintenanceWindow  string
	ReplicationEnabled bool
	// Additional parameters:
	CacheSize      string
	LogDirectory   string
	Timezone       string
	SSLEnabled     bool
	Locale         string
	Autovacuum     bool
	DataEncryption bool
}

// --- Implementation of types.Service ---

func (s *PostgreSQLService) Name() string {
	return "postgresql"
}

func (s *PostgreSQLService) Version() string {
	return s.VersionField
}

func (s *PostgreSQLService) RequiredInterfaces() []string {
	return []string{"Installer", "Hooker", "Configurator", "Updater", "Uninstaller", "Validator", "PackageInfoProvider"}
}

//////////////////////////////
// PackageInfoProvider Methods
//////////////////////////////

// GetPackageName returns the package name based on the detected package manager.
func (s *PostgreSQLService) GetPackageName(pm string) string {
	pm = strings.ToLower(pm)
	switch pm {
	case "apt":
		return "postgresql"
	case "dnf", "yum":
		return "postgresql-server"
		// You may support custom keys as needed.
	default:
		return "postgresql"
	}
}

// GetPackageVersion returns the PostgreSQL version.
func (s *PostgreSQLService) GetPackageVersion() string {
	return s.VersionField
}

// GetRepository returns the repository address for the given package manager.
// The service defines its repository at the service level.
func (s *PostgreSQLService) GetRepository(pm string) string {
	pm = strings.ToLower(pm)
	switch pm {
	case "apt":
		// Return the deb repository string.
		return "deb http://apt.postgresql.org/pub/repos/apt/ $(lsb_release -cs)-pgdg main"
	case "dnf", "yum":
		// Return the RPM repository URL.
		return "https://download.postgresql.org/pub/repos/yum/reporpms/EL-$(rpm -E %{rhel})-x86_64/"
	default:
		return ""
	}
}

//////////////////////////////
// Installer Methods (Sample)
//////////////////////////////

func (s *PostgreSQLService) PreInstall(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Debug("PostgreSQL PreInstall: checking prerequisites")
	return nil
}

func (s *PostgreSQLService) Install(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("Starting PostgreSQL installation", zap.String("version", s.VersionField))
	v, err := semver.ParseTolerant(s.VersionField)
	if err != nil {
		return err
	}
	if v.Major == 1 {
		return s.installLegacy(ctx, svcCtx)
	}
	return s.installModern(ctx, svcCtx)
}

func (s *PostgreSQLService) DryRunInstall(ctx context.Context, svcCtx *types.ServiceContext) ([]string, error) {
	steps := []string{
		"Check prerequisites",
		"Download PostgreSQL package",
		"Verify checksum",
		"Configure installation paths",
	}
	return steps, nil
}

func (s *PostgreSQLService) PostInstall(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("Post-install: PostgreSQL installation complete")
	return nil
}

func (s *PostgreSQLService) installLegacy(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("Executing legacy installation steps for PostgreSQL")
	time.Sleep(1 * time.Second)
	return nil
}

func (s *PostgreSQLService) installModern(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("Executing modern installation steps for PostgreSQL")
	time.Sleep(1 * time.Second)
	return nil
}

// (Implement Runner, Configurator, Updater, Uninstaller, Validator methods as needed...)
