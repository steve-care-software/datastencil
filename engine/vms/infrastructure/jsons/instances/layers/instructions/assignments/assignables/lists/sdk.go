package lists

import (
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists"
	json_fetches "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/lists/fetches"
)

// NewAdapter creates a new adapter
func NewAdapter() lists.Adapter {
	fetchAdapter := json_fetches.NewAdapter()
	builder := lists.NewBuilder()
	return createAdapter(
		fetchAdapter.(*json_fetches.Adapter),
		builder,
	)
}
