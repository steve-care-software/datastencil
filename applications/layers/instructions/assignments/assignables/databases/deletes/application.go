package deletes

import (
	"github.com/steve-care-software/datastencil/applications/layers/instructions/failures"
	databases_deletes "github.com/steve-care-software/datastencil/domain/instances/databases/commits/actions/modifications/deletes"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions/assignments/assignables/databases/deletes"
	"github.com/steve-care-software/datastencil/domain/stacks"
)

type application struct {
	assignableBuilder stacks.AssignableBuilder
	deleteBuilder     databases_deletes.Builder
}

func createApplication(
	assignableBuilder stacks.AssignableBuilder,
	deleteBuilder databases_deletes.Builder,
) Application {
	out := application{
		assignableBuilder: assignableBuilder,
		deleteBuilder:     deleteBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable deletes.Delete) (stacks.Assignable, *uint, error) {
	indexVar := assignable.Index()
	pIndex, err := frame.FetchUnsignedInt(indexVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, nil
	}

	lengthVar := assignable.Length()
	pLength, err := frame.FetchUnsignedInt(lengthVar)
	if err != nil {
		code := failures.CouldNotFetchUnsignedIntegerFromFrame
		return nil, &code, nil
	}

	delete, err := app.deleteBuilder.Create().WithIndex(*pIndex).WithLength(*pLength).Now()
	if err != nil {
		return nil, nil, err
	}

	ins, err := app.assignableBuilder.Create().WithDelete(delete).Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}