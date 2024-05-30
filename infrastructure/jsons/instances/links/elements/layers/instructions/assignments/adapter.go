package assignments

import (
	"encoding/json"

	"github.com/steve-care-software/datastencil/domain/instances/pointers/resources/logics/bridges/layers/instructions/assignments"
	json_assignables "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/assignments/assignables"
)

// Adapter represents the adapter
type Adapter struct {
	assignableAdapter *json_assignables.Adapter
	builder           assignments.Builder
}

func createAdapter(
	assignableAdapter *json_assignables.Adapter,
	builder assignments.Builder,
) assignments.Adapter {
	out := Adapter{
		assignableAdapter: assignableAdapter,
		builder:           builder,
	}

	return &out
}

// ToBytes converts instance to bytes
func (app *Adapter) ToBytes(ins assignments.Assignment) ([]byte, error) {
	ptr, err := app.AssignmentToStruct(ins)
	if err != nil {
		return nil, err
	}

	js, err := json.Marshal(ptr)
	if err != nil {
		return nil, err
	}

	return js, nil
}

// ToInstance converts bytes to instance
func (app *Adapter) ToInstance(bytes []byte) (assignments.Assignment, error) {
	ins := new(Assignment)
	err := json.Unmarshal(bytes, ins)
	if err != nil {
		return nil, err
	}

	return app.StructToAssignment(*ins)
}

// AssignmentToStruct converts an ssignment to struct
func (app *Adapter) AssignmentToStruct(ins assignments.Assignment) (*Assignment, error) {
	ptr, err := app.assignableAdapter.AssignableToStruct(ins.Assignable())
	if err != nil {
		return nil, err
	}

	return &Assignment{
		Name:       ins.Name(),
		Assignable: *ptr,
	}, nil
}

// StructToAssignment converts a struct to assignment
func (app *Adapter) StructToAssignment(str Assignment) (assignments.Assignment, error) {
	ins, err := app.assignableAdapter.StructToAssignable(str.Assignable)
	if err != nil {
		return nil, err
	}

	return app.builder.Create().
		WithName(str.Name).
		WithAssignable(ins).
		Now()
}
