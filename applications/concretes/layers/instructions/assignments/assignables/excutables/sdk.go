package executables

import (
	"github.com/steve-care-software/datastencil/applications"
	instruction_executables "github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executables"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application for tests
func NewApplication(
	localAppBuilder applications.LocalBuilder,
	remoteAppBuilder applications.RemoteBuilder,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	return createApplication(
		localAppBuilder,
		remoteAppBuilder,
		assignableBuilder,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable instruction_executables.Executable) (stacks.Assignable, *uint, error)
}
