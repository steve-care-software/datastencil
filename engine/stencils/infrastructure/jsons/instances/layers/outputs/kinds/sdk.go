package kinds

import "github.com/steve-care-software/webx/engine/stencils/domain/instances/layers/outputs/kinds"

// NewAdapter creates a new adapter
func NewAdapter() kinds.Adapter {
	builder := kinds.NewBuilder()
	return createAdapter(
		builder,
	)
}