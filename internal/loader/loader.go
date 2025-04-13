// internal/loader/loader.go
package loader

import (
	"errors"
	"path/filepath"
	"plugin"

	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
)

// LoadService loads a plugin module given its file path.
// It expects the module to export a `NewService` symbol with signature: func() types.Service.
func LoadService(path string) (types.Service, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	p, err := plugin.Open(absPath)
	if err != nil {
		return nil, err
	}

	sym, err := p.Lookup("NewService")
	if err != nil {
		return nil, err
	}

	newServiceFunc, ok := sym.(func() types.Service)
	if !ok {
		return nil, errors.New("invalid plugin signature: expected func() types.Service")
	}

	return newServiceFunc(), nil
}
