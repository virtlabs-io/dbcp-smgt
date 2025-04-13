package unit

import (
	"os"
	"testing"

	"github.com/virtlabs-io/dbcp-smgt/internal/distro"
)

func TestDetectPackageManager(t *testing.T) {
	// Note: this test depends on your current system.
	pm := distro.DetectPackageManager()
	if pm == distro.PackageManagerUNKNOWN {
		t.Errorf("Expected a known package manager, got %s", pm)
	} else {
		t.Logf("Detected package manager: %s", pm)
	}
}

// Alternatively, you could create a helper that loads a test os-release file.
func TestParseOsRelease(t *testing.T) {
	// Create a temporary os-release file with sample content.
	content := `
ID="ubuntu"
ID_LIKE="debian"
`
	tmpFile := "os-release.test"
	err := os.WriteFile(tmpFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(tmpFile)

	// Temporarily override the file path by modifying the function or adding a new test helper.
	// Here we assume you might have refactored detection into a helper that accepts a file path.
	// For illustration purposes only:
	// pm, err := distro.DetectPackageManagerFrom(tmpFile)
	// if err != nil {
	//     t.Fatalf("Detection failed: %v", err)
	// }
	// if pm != distro.PackageManagerAPT {
	//     t.Errorf("Expected APT, got %s", pm)
	// }
}
