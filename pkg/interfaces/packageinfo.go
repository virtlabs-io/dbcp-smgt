// pkg/interfaces/packageinfo.go
package interfaces

// PackageInfoProvider lets a service supply package metadata.
type PackageInfoProvider interface {
	// GetPackageName returns the package name to install for the given package manager.
	GetPackageName(pm string) string
	// GetPackageVersion returns the package version as a string.
	GetPackageVersion() string
	// GetRepository returns the repository specification for the given package manager.
	// This may be a deb-repo string for APT, an rpm-repo definition for YUM/DNF, or even a custom repo URL.
	GetRepository(pm string) string
}
