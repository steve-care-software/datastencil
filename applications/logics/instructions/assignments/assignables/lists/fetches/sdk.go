package fetches

import (
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication() Application {
	return createApplication()
}

// Application represents an application
type Application interface {
	Execute(frame stacks.Frame, assignable fetches.Fetch) (stacks.Assignable, *uint, error)
}
