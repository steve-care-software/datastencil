package assignments

import (
	"github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execAssignableApp assignables.Application,
) Application {
	assignmentBuilder := stacks.NewAssignmentBuilder()
	return createApplication(
		execAssignableApp,
		assignmentBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignment assignments.Assignment) (stacks.Assignment, *uint, error)
}
