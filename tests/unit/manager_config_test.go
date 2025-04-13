// tests/unit/manager_config_test.go
package unit

import (
	"context"
	"testing"

	"go.uber.org/zap/zaptest"

	"github.com/virtlabs-io/dbcp-smgt/internal/core"
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
	"github.com/virtlabs-io/dbcp-smgt/tests/unit/fakes"
)

func TestExecuteConfigure(t *testing.T) {
	logger := zaptest.NewLogger(t)
	svcCtx := &types.ServiceContext{
		Version: "1.0.0",
		Config:  map[string]interface{}{"env": "test"},
		Logger:  logger,
		DryRun:  false,
	}

	manager := core.NewServiceManager(svcCtx)

	// Use the fake service from the fakes package that implements Configurator.
	fakeSvc := &fakes.FakeService{} // Ensure FakeService also implements Configurator (you may need to add stub methods)

	if err := manager.ExecuteConfigure(context.Background(), fakeSvc); err != nil {
		t.Fatalf("Expected no error on ExecuteConfigure, but got: %v", err)
	}
}
