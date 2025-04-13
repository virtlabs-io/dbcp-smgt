// tests/unit/fakes/fake_service.go
package fakes

import (
	"context"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// FakeService implements Installer, Hooker, and Configurator interfaces along with types.Service.
type FakeService struct {
	types.BaseService // Embeds default (no-op) hook implementations.
}

// --- Implementation of types.Service ---

// Name returns the service name.
func (f *FakeService) Name() string {
	return "FakeService"
}

// Version returns the service version.
func (f *FakeService) Version() string {
	return "1.0.0"
}

// RequiredInterfaces returns the list of interface names this service implements.
func (f *FakeService) RequiredInterfaces() []string {
	return []string{"Installer", "Hooker", "Configurator"}
}

// --- Implementation of Installer interface ---

func (f *FakeService) PreInstall(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("FakeService: PreInstall")
	return nil
}

func (f *FakeService) Install(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("FakeService: Install")
	return nil
}

func (f *FakeService) DryRunInstall(ctx context.Context, svcCtx *types.ServiceContext) ([]string, error) {
	steps := []string{
		"FakeService: DryRunInstall step 1",
		"FakeService: DryRunInstall step 2",
	}
	return steps, nil
}

func (f *FakeService) PostInstall(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("FakeService: PostInstall")
	return nil
}

// --- Implementation of Hooker interface ---

func (f *FakeService) PostStart(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("FakeService: PostStart")
	return nil
}

// --- Implementation of Configurator interface ---

func (f *FakeService) PreConfigure(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("FakeService: PreConfigure")
	return nil
}

func (f *FakeService) Configure(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("FakeService: Configure")
	return nil
}

func (f *FakeService) PostConfigure(ctx context.Context, svcCtx *types.ServiceContext) error {
	svcCtx.Logger.Info("FakeService: PostConfigure")
	return nil
}
