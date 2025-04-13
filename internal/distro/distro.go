// internal/distro/distro.go
package distro

import (
	"bufio"
	"os"
	"strings"
)

// PackageManager represents a type for Linux package managers.
type PackageManager string

const (
	PackageManagerAPT     PackageManager = "apt" // Debian, Ubuntu, etc.
	PackageManagerDNF     PackageManager = "dnf" // Fedora, RHEL 8+, etc.
	PackageManagerYUM     PackageManager = "yum" // Older RHEL/CentOS/Oracle Linux.
	PackageManagerUNKNOWN PackageManager = "unknown"
)

// DetectPackageManager reads /etc/os-release and returns the package manager.
func DetectPackageManager() PackageManager {
	file, err := os.Open("/etc/os-release")
	if err != nil {
		return PackageManagerUNKNOWN
	}
	defer file.Close()

	var id, idLike string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "ID=") {
			parts := strings.SplitN(line, "=", 2)
			id = strings.Trim(parts[1], "\"")
		} else if strings.HasPrefix(line, "ID_LIKE=") {
			parts := strings.SplitN(line, "=", 2)
			idLike = strings.Trim(parts[1], "\"")
		}
	}
	id = strings.ToLower(id)
	idLike = strings.ToLower(idLike)

	// Look for Debian/Ubuntu derivatives.
	if id == "ubuntu" || id == "debian" || strings.Contains(idLike, "debian") {
		return PackageManagerAPT
	}
	// Look for Fedora, RHEL, CentOS, etc.
	if id == "fedora" || id == "rhel" || id == "centos" ||
		strings.Contains(idLike, "rhel") ||
		strings.Contains(idLike, "fedora") ||
		strings.Contains(idLike, "centos") {
		// Assume modern distributions use DNF.
		return PackageManagerDNF
	}
	return PackageManagerUNKNOWN
}
