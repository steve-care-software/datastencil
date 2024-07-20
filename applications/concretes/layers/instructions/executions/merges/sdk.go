package merges

import (
	"github.com/steve-care-software/datastencil/applications"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/executions/merges"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, executable applications.Application, assignment merges.Merge) (*uint, error)
}
