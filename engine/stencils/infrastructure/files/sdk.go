package files

import (
	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/contexts"
)

const readWritePermissionBits = 0755

// NewContextRepository creates a new context repository
func NewContextRepository(
	adapter contexts.Adapter,
	endPath []string,
) contexts.Repository {
	hashAdapter := hash.NewAdapter()
	return createContextRepository(
		adapter,
		hashAdapter,
		endPath,
	)
}

// NewContextService creates a new context service
func NewContextService(
	adapter contexts.Adapter,
	endPath []string,
) contexts.Service {
	hashAdapter := hash.NewAdapter()
	return createContextService(
		adapter,
		hashAdapter,
		endPath,
	)
}
