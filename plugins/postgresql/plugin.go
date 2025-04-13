// services/postgresql/plugin.go
package main

import (
	"github.com/virtlabs-io/dbcp-smgt/pkg/types"
	"github.com/virtlabs-io/dbcp-smgt/services/postgresql"
)

// NewService is the exported symbol for the plugin.
// It returns a new instance of PostgreSQLService that implements the types.Service interface.
func NewService() types.Service {
	return &postgresql.PostgreSQLService{
		// You can initialize defaults or leave these to be configured later.
		VersionField: "13.3",
		// Other fields may be initialized here or loaded via module config.
	}
}
