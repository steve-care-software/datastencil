package cryptography

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

// NewApplication creates a new application
func NewApplication(
	execDecryptApp decrypts.Application,
	execEncryptApp encrypts.Application,
) Application {
	return createApplication(
		execDecryptApp,
		execEncryptApp,
	)
}

// Application represents an execution account application
type Application interface {
	Execute(frame stacks.Frame, assignable cryptography.Cryptography) (stacks.Assignable, *uint, error)
}
