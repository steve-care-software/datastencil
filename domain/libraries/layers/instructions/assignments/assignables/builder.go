package assignables

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/accounts"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/cryptography"
)

type builder struct {
	hashAdapter hash.Adapter
	bytes       bytes.Bytes
	constant    constants.Constant
	account     accounts.Account
	crypto      cryptography.Cryptography
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		constant:    nil,
		account:     nil,
		crypto:      nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder(
		app.hashAdapter,
	)
}

// WithBytes add bytes to the builder
func (app *builder) WithBytes(bytes bytes.Bytes) Builder {
	app.bytes = bytes
	return app
}

// WithConsant adds a constant to the builder
func (app *builder) WithConsant(constant constants.Constant) Builder {
	app.constant = constant
	return app
}

// WithAccount adds an account to the builder
func (app *builder) WithAccount(account accounts.Account) Builder {
	app.account = account
	return app
}

// WithCryptography adds a cryptography to the builder
func (app *builder) WithCryptography(cryptography cryptography.Cryptography) Builder {
	app.crypto = cryptography
	return app
}

// Now builds a new Assignable instance
func (app *builder) Now() (Assignable, error) {
	data := [][]byte{}
	if app.bytes != nil {
		data = append(data, []byte("bytes"))
		data = append(data, app.bytes.Hash().Bytes())
	}

	if app.constant != nil {
		data = append(data, []byte("constant"))
		data = append(data, app.constant.Hash().Bytes())
	}

	if app.account != nil {
		data = append(data, []byte("account"))
		data = append(data, app.account.Hash().Bytes())
	}

	if app.crypto != nil {
		data = append(data, []byte("crypto"))
		data = append(data, app.crypto.Hash().Bytes())
	}

	if len(data) <= 0 {
		return nil, errors.New("the Assignable is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(data)
	if err != nil {
		return nil, err
	}

	if app.bytes != nil {
		return createAssignableWithBytes(*pHash, app.bytes), nil
	}

	if app.constant != nil {
		return createAssignableWithConstant(*pHash, app.constant), nil
	}

	return createAssignableWithAccount(*pHash, app.account), nil
}