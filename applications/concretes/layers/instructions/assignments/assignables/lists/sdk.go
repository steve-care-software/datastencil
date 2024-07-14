package lists

import (
	application_fetches "github.com/steve-care-software/datastencil/applications/concretes/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/lists"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	fetchApplication application_fetches.Application,
) Application {
	assignableBuilder := stacks.NewAssignableBuilder()
	assignablesBuilder := stacks.NewAssignablesBuilder()
	return createApplication(
		fetchApplication,
		assignableBuilder,
		assignablesBuilder,
	)
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable lists.List) (stacks.Assignable, *uint, error)
}