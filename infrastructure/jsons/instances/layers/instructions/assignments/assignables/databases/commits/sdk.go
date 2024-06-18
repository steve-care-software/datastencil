package commits

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/databases/commits"
)

// NewAdapter creates a new adapter
func NewAdapter() commits.Adapter {
	builder := commits.NewBuilder()
	return createAdapter(
		builder,
	)
}