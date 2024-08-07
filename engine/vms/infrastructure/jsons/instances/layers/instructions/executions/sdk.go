package executions

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/executions"
	json_merges "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/executions/merges"
)

// NewAdapter creates a new adapter
func NewAdapter() executions.Adapter {
	mergeAdapter := json_merges.NewAdapter()
	builder := executions.NewBuilder()
	contentBuilder := executions.NewContentBuilder()
	return createAdapter(
		mergeAdapter.(*json_merges.Adapter),
		builder,
		contentBuilder,
	)
}
