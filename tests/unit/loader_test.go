package unit

import (
	"path/filepath"
	"testing"

	"github.com/virtlabs-io/dbcp-smgt/internal/loader"
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

func TestLoadService_PostgreSQL(t *testing.T) {
	// Compute absolute path for the plugin built in the bin/ directory.
	pluginPath, err := filepath.Abs("../../bin/postgresql.so")
	if err != nil {
		t.Fatalf("Failed to get absolute path: %v", err)
	}

	svc, err := loader.LoadService(pluginPath)
	if err != nil {
		t.Fatalf("Error loading PostgreSQL plugin: %v", err)
	}
	if svc == nil {
		t.Fatal("Expected a valid service, got nil")
	}
	// Verify that the loaded service satisfies the types.Service interface.
	var _ types.Service = svc
}
