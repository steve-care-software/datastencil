package retrieves

import (
	"errors"
	"fmt"

	"github.com/steve-care-software/datastencil/domain/instances/databases"
	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments/assignables/databases/retrieves"
	"github.com/steve-care-software/datastencil/domain/stacks"
	"github.com/steve-care-software/datastencil/domain/stacks/failures"
)

type application struct {
	repository         databases.Repository
	assignableBuilder  stacks.AssignableBuilder
	assignablesBuilder stacks.AssignablesBuilder
}

func createApplication(
	repository databases.Repository,
	assignableBuilder stacks.AssignableBuilder,
	assignablesBuilder stacks.AssignablesBuilder,
) Application {
	out := application{
		repository:         repository,
		assignableBuilder:  assignableBuilder,
		assignablesBuilder: assignablesBuilder,
	}

	return &out
}

// Execute executes the application
func (app *application) Execute(frame stacks.Frame, assignable retrieves.Retrieve) (stacks.Assignable, *uint, error) {
	builder := app.assignableBuilder.Create()
	if assignable.IsList() {
		pathList, err := app.repository.List()
		if err != nil {
			code := failures.CouldNotRetrieveListFromRepository
			return nil, &code, err
		}

		if len(pathList) <= 0 {
			code := failures.CouldNotRetrieveEmptyListFromRepository
			return nil, &code, errors.New("the path list returned from the database repository's List method is empty")
		}

		pathAssignableList := []stacks.Assignable{}
		for _, onePath := range pathList {
			assList := []stacks.Assignable{}
			for _, oneStr := range onePath {
				assignable, err := app.assignableBuilder.Create().
					WithString(oneStr).
					Now()

				if err != nil {
					return nil, nil, err
				}

				assList = append(assList, assignable)
			}

			assignables, err := app.assignablesBuilder.Create().
				WithList(assList).
				Now()

			if err != nil {
				return nil, nil, err
			}

			pathAssignable, err := app.assignableBuilder.Create().
				WithList(assignables).
				Now()

			if err != nil {
				return nil, nil, err
			}

			pathAssignableList = append(pathAssignableList, pathAssignable)
		}

		assignables, err := app.assignablesBuilder.Create().
			WithList(pathAssignableList).
			Now()

		if err != nil {
			return nil, nil, err
		}

		builder.WithList(assignables)
	}

	if assignable.IsExists() {
		pathVar := assignable.Exists()
		pathAssignables, err := frame.FetchList(pathVar)
		if err != nil {
			code := failures.CouldNotFetchListFromFrame
			return nil, &code, err
		}

		pathValues := []string{}
		pathList := pathAssignables.List()
		for _, oneAssignable := range pathList {
			if !oneAssignable.IsString() {
				code := failures.CouldNotFetchStringFromList
				str := fmt.Sprintf("the path (name: %s) contains invalid elements", pathVar)
				return nil, &code, errors.New(str)
			}

			pathValues = append(pathValues, *oneAssignable.String())
		}

		pExists, err := app.repository.Exists(pathValues)
		if err != nil {
			code := failures.CouldNotExecuteExistsFromRepository
			return nil, &code, err
		}

		builder.WithBool(*pExists)
	}

	if assignable.IsRetrieve() {
		pathVar := assignable.Retrieve()
		pathAssignables, err := frame.FetchList(pathVar)
		if err != nil {
			code := failures.CouldNotFetchListFromFrame
			return nil, &code, err
		}

		pathValues := []string{}
		pathList := pathAssignables.List()
		for _, oneAssignable := range pathList {
			if !oneAssignable.IsString() {
				code := failures.CouldNotFetchStringFromList
				str := fmt.Sprintf("the path (name: %s) contains invalid elements", pathVar)
				return nil, &code, errors.New(str)
			}

			pathValues = append(pathValues, *oneAssignable.String())
		}

		retDatabase, err := app.repository.Retrieve(pathValues)
		if err != nil {
			code := failures.CouldNotRetrieveFromRepository
			return nil, &code, err
		}

		builder.WithDatabase(retDatabase)
	}

	ins, err := builder.Now()
	if err != nil {
		return nil, nil, err
	}

	return ins, nil, nil
}
