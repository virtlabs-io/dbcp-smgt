// internal/pkgmgr/pkgmgr.go
package pkgmgr

import (
	"fmt"
	"os/exec"

	"github.com/virtlabs-io/dbcp-smgt/internal/distro"
)

// InstallOptions holds details for the package installation.
type InstallOptions struct {
	PackageName    string
	PackageVersion string
	Repository     string
	PackageManager distro.PackageManager
}

// AddRepository adds the repository via the appropriate command for the package manager.
func AddRepository(opts InstallOptions) error {
	if opts.Repository == "" {
		// No repository specification; nothing to do.
		return nil
	}
	switch opts.PackageManager {
	case distro.PackageManagerAPT:
		// For example, use add-apt-repository (this may require root privileges).
		cmd := exec.Command("add-apt-repository", "-y", opts.Repository)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to add APT repository: %v, output: %s", err, output)
		}
	case distro.PackageManagerDNF, distro.PackageManagerYUM:
		// For RPM-based systems, one way is to create a repo file.
		cmd := exec.Command("sh", "-c", fmt.Sprintf("echo '[custom-repo]\nname=Custom Repo\nbaseurl=%s\nenabled=1' > /etc/yum.repos.d/custom.repo", opts.Repository))
		output, err := cmd.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to add RPM repository: %v, output: %s", err, output)
		}
	default:
		return fmt.Errorf("unsupported package manager for adding repository: %s", opts.PackageManager)
	}
	return nil
}

// InstallPackage installs the package using the proper package manager.
func InstallPackage(opts InstallOptions) error {
	var cmd *exec.Cmd
	pkgFull := opts.PackageName
	// For APT, you can specify version by "package=version"
	if opts.PackageManager == distro.PackageManagerAPT && opts.PackageVersion != "" {
		pkgFull = fmt.Sprintf("%s=%s", opts.PackageName, opts.PackageVersion)
	}
	switch opts.PackageManager {
	case distro.PackageManagerAPT:
		// Update and install.
		updateCmd := exec.Command("apt-get", "update")
		if out, err := updateCmd.CombinedOutput(); err != nil {
			return fmt.Errorf("apt-get update error: %v, output: %s", err, out)
		}
		cmd = exec.Command("apt-get", "install", "-y", pkgFull)
	case distro.PackageManagerDNF:
		cmd = exec.Command("dnf", "install", "-y", pkgFull)
	case distro.PackageManagerYUM:
		cmd = exec.Command("yum", "install", "-y", pkgFull)
	default:
		return fmt.Errorf("unsupported package manager: %s", opts.PackageManager)
	}
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to install package: %v, output: %s", err, output)
	}
	return nil
}
