package keys

import (
	"github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	"github.com/steve-care-software/datastencil/stencils/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	"github.com/steve-care-software/datastencil/stencils/domain/instances/layers/instructions/assignments/assignables/cryptography/keys"
	"github.com/steve-care-software/datastencil/stencils/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	encApp encryptions.Application,
	sigApp signatures.Application,
) Application {
	return createApplication(
		encApp,
		sigApp,
	)
}

// Application represents a key application
type Application interface {
	Execute(frame stacks.Frame, assignable keys.Key) (stacks.Assignable, *uint, error)
}
