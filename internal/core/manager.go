package core

import (
	"context"
	"errors"

	"go.uber.org/zap"

	"github.com/virtlabs-io/dbcp-smgt/internal/distro"
	"github.com/virtlabs-io/dbcp-smgt/internal/pkgmgr"
	"github.com/virtlabs-io/dbcp-smgt/pkg/interfaces"
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// ServiceManager orchestrates service lifecycle operations.
type ServiceManager struct {
	Context *types.ServiceContext
}

// NewServiceManager creates a new ServiceManager using the provided ServiceContext.
func NewServiceManager(ctx *types.ServiceContext) *ServiceManager {
	return &ServiceManager{Context: ctx}
}

// ExecuteInstall runs the installation lifecycle phase.
// It leverages PackageInfoProvider if available.
func (m *ServiceManager) ExecuteInstall(ctx context.Context, service types.Service) error {
	// Detect current package manager.
	pm := distro.DetectPackageManager()
	m.Context.Logger.Info("Detected package manager", zap.String("package_manager", string(pm)))

	// Execute PreInstall hook if provided.
	if hooker, ok := service.(interfaces.Hooker); ok {
		if err := hooker.PreInstall(ctx, m.Context); err != nil {
			return err
		}
	}

	// If the service implements PackageInfoProvider, use it.
	if pkgProvider, ok := service.(interfaces.PackageInfoProvider); ok {
		opts := pkgmgr.InstallOptions{
			PackageName:    pkgProvider.GetPackageName(string(pm)),
			PackageVersion: pkgProvider.GetPackageVersion(),
			Repository:     pkgProvider.GetRepository(string(pm)),
			PackageManager: pm,
		}
		m.Context.Logger.Info("Installing package",
			zap.String("pkgName", opts.PackageName),
			zap.String("pkgVersion", opts.PackageVersion),
			zap.String("repository", opts.Repository),
			zap.String("pkgManager", string(opts.PackageManager)),
		)
		if m.Context.DryRun {
			m.Context.Logger.Info("Dry-run mode: skipping actual repository addition and installation", zap.Any("opts", opts))
			return nil
		}
		if err := pkgmgr.AddRepository(opts); err != nil {
			return err
		}
		return pkgmgr.InstallPackage(opts)
	} else if installer, ok := service.(interfaces.Installer); ok {
		// Fallback: if PackageInfoProvider is not implemented, use service's own Installer.
		if m.Context.DryRun {
			steps, _ := installer.DryRunInstall(ctx, m.Context)
			m.Context.Logger.Info("Dry-run installation steps", zap.Strings("steps", steps))
			return nil
		}
		return installer.Install(ctx, m.Context)
	}
	return errors.New("Installer interface not implemented")
}

// ExecuteConfigure runs the configuration lifecycle phase.
func (m *ServiceManager) ExecuteConfigure(ctx context.Context, service types.Service) error {
	if configurator, ok := service.(interfaces.Configurator); ok {
		if err := configurator.PreConfigure(ctx, m.Context); err != nil {
			return err
		}
		if err := configurator.Configure(ctx, m.Context); err != nil {
			return err
		}
		return configurator.PostConfigure(ctx, m.Context)
	}
	return errors.New("Configurator interface not implemented")
}
