// tests/unit/manager_test.go
package unit

import (
	"context"
	"testing"

	"go.uber.org/zap/zaptest"

	"github.com/virtlabs-io/dbcp-smgt/internal/core"
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
	"github.com/virtlabs-io/dbcp-smgt/tests/unit/fakes"
)

func TestExecuteInstall(t *testing.T) {
	// Create a test logger using zaptest.
	logger := zaptest.NewLogger(t)

	// Create a ServiceContext with test configuration.
	svcCtx := &types.ServiceContext{
		Version: "1.0.0",
		Config: map[string]interface{}{
			"env": "test",
		},
		Logger: logger,
		DryRun: false,
	}

	// Instantiate the ServiceManager.
	manager := core.NewServiceManager(svcCtx)

	// Create an instance of FakeService.
	fakeSvc := &fakes.FakeService{}

	// Call ExecuteInstall and verify no error is returned.
	if err := manager.ExecuteInstall(context.Background(), fakeSvc); err != nil {
		t.Fatalf("Expected no error from ExecuteInstall, but got: %v", err)
	}
}
