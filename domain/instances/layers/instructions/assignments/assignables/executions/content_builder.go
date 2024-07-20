package executions

import (
	"errors"

	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/amounts"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/begins"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/executes"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/heads"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/inits"
	"github.com/steve-care-software/datastencil/domain/instances/layers/instructions/assignments/assignables/executions/retrieves"
	"github.com/steve-care-software/historydb/domain/hash"
)

type contentBuilder struct {
	hashAdapter hash.Adapter
	isList      bool
	init        inits.Init
	begin       begins.Begin
	execute     executes.Execute
	retrieve    retrieves.Retrieve
	amount      amounts.Amount
	head        heads.Head
}

func createContentBuilder(
	hashAdapter hash.Adapter,
) ContentBuilder {
	out := contentBuilder{
		hashAdapter: hashAdapter,
		isList:      false,
		init:        nil,
		begin:       nil,
		execute:     nil,
		retrieve:    nil,
		amount:      nil,
		head:        nil,
	}

	return &out
}

// Create initializes the contentBuilder
func (app *contentBuilder) Create() ContentBuilder {
	return createContentBuilder(
		app.hashAdapter,
	)
}

// WithInit adds an init to the builder
func (app *contentBuilder) WithInit(init inits.Init) ContentBuilder {
	app.init = init
	return app
}

// WithBegin adds a begin to the builder
func (app *contentBuilder) WithBegin(begin begins.Begin) ContentBuilder {
	app.begin = begin
	return app
}

// WithExecute adds an execute to the builder
func (app *contentBuilder) WithExecute(execute executes.Execute) ContentBuilder {
	app.execute = execute
	return app
}

// WithRetrieve adds a retrieve to the builder
func (app *contentBuilder) WithRetrieve(retrieve retrieves.Retrieve) ContentBuilder {
	app.retrieve = retrieve
	return app
}

// WithAmount adds an amount to the builder
func (app *contentBuilder) WithAmount(amount amounts.Amount) ContentBuilder {
	app.amount = amount
	return app
}

// WithHead adds an head to the builder
func (app *contentBuilder) WithHead(head heads.Head) ContentBuilder {
	app.head = head
	return app
}

// IsList flags the builder as a list
func (app *contentBuilder) IsList() ContentBuilder {
	app.isList = true
	return app
}

// Now builds a new Execution instance
func (app *contentBuilder) Now() (Content, error) {
	bytes := [][]byte{}
	if app.isList {
		bytes = append(bytes, []byte("isList"))
	}

	if app.init != nil {
		bytes = append(bytes, []byte("init"))
		bytes = append(bytes, app.init.Hash().Bytes())
	}

	if app.begin != nil {
		bytes = append(bytes, []byte("begin"))
		bytes = append(bytes, app.begin.Hash().Bytes())
	}

	if app.execute != nil {
		bytes = append(bytes, []byte("execute"))
		bytes = append(bytes, app.execute.Hash().Bytes())
	}

	if app.retrieve != nil {
		bytes = append(bytes, []byte("retrieve"))
		bytes = append(bytes, app.retrieve.Hash().Bytes())
	}

	if app.amount != nil {
		bytes = append(bytes, []byte("amount"))
		bytes = append(bytes, app.amount.Hash().Bytes())
	}

	if app.head != nil {
		bytes = append(bytes, []byte("head"))
		bytes = append(bytes, app.head.Hash().Bytes())
	}

	amount := len(bytes)
	if amount != 1 && amount != 2 {
		return nil, errors.New("the Execution is invalid")
	}

	pHash, err := app.hashAdapter.FromMultiBytes(bytes)
	if err != nil {
		return nil, err
	}

	if app.isList {
		return createContentWithList(*pHash), nil
	}

	if app.init != nil {
		return createContentWithInit(*pHash, app.init), nil
	}

	if app.begin != nil {
		return createContentWithBegin(*pHash, app.begin), nil
	}

	if app.execute != nil {
		return createContentWithExecute(*pHash, app.execute), nil
	}

	if app.retrieve != nil {
		return createContentWithRetrieve(*pHash, app.retrieve), nil
	}

	if app.amount != nil {
		return createContentWithAmount(*pHash, app.amount), nil
	}

	return createContentWithHead(*pHash, app.head), nil
}
