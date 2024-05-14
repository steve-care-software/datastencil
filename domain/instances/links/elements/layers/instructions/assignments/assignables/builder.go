package assignables

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/lists"
)

type builder struct {
	hashAdapter hash.Adapter
	bytes       bytes.Bytes
	constant    constants.Constant
	crypto      cryptography.Cryptography
	compiler    compilers.Compiler
	database    databases.Database
	list        lists.List
}

func createBuilder(
	hashAdapter hash.Adapter,
) Builder {
	out := builder{
		hashAdapter: hashAdapter,
		bytes:       nil,
		constant:    nil,
		crypto:      nil,
		compiler:    nil,
		database:    nil,
		list:        nil,
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

// WithCryptography adds a cryptography to the builder
func (app *builder) WithCryptography(cryptography cryptography.Cryptography) Builder {
	app.crypto = cryptography
	return app
}

// WithCompiler adds a compiler to the builder
func (app *builder) WithCompiler(compiler compilers.Compiler) Builder {
	app.compiler = compiler
	return app
}

// WithDatabase adds a database to the builder
func (app *builder) WithDatabase(database databases.Database) Builder {
	app.database = database
	return app
}

// WithList adds a list to the builder
func (app *builder) WithList(list lists.List) Builder {
	app.list = list
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

	if app.crypto != nil {
		data = append(data, []byte("crypto"))
		data = append(data, app.crypto.Hash().Bytes())
	}

	if app.compiler != nil {
		data = append(data, []byte("compiler"))
		data = append(data, app.compiler.Hash().Bytes())
	}

	if app.database != nil {
		data = append(data, []byte("database"))
		data = append(data, app.database.Hash().Bytes())
	}

	if app.list != nil {
		data = append(data, []byte("list"))
		data = append(data, app.list.Hash().Bytes())
	}

	if len(data) != 2 {
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

	if app.crypto != nil {
		return createAssignableWithCryptography(*pHash, app.crypto), nil
	}

	if app.database != nil {
		return createAssignableWithDatabase(*pHash, app.database), nil
	}

	if app.list != nil {
		return createAssignableWithList(*pHash, app.list), nil
	}

	return createAssignableWithCompiler(*pHash, app.compiler), nil
}
