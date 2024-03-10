package encrypts

import (
	"github.com/steve-care-software/datastencil/domain/encryptors"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/cryptography/encrypts"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	encryptor         encryptors.Encryptor
	assignableBuilder stacks.AssignableBuilder
}

func createApplication(
	encryptor encryptors.Encryptor,
	assignableBuilder stacks.AssignableBuilder,
) Application {
	out := application{
		encryptor:         encryptor,
		assignableBuilder: assignableBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable encrypts.Encrypt) (stacks.Assignable, error) {
	msgVar := assignable.Message()
	message, err := frame.FetchBytes(msgVar)
	if err != nil {
		return nil, err
	}

	passVar := assignable.Password()
	password, err := frame.FetchBytes(passVar)
	if err != nil {
		return nil, err
	}

	cipher, err := app.encryptor.Encrypt(message, password)
	if err != nil {
		return nil, err
	}

	return app.assignableBuilder.Create().
		WithBytes(cipher).
		Now()
}