package lists

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/lists/deletes"
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/lists/inserts"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/lists"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	insertApp inserts.Application,
	deleteApp deletes.Application,
) Application {
	return createApplication(
		insertApp,
		deleteApp,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, list lists.List) (stacks.Assignment, *uint, error)
}
