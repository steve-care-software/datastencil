package pointers

import "errors"

type builder struct {
	list []Pointer
}

func createBuilder() Builder {
	out := builder{
		list: nil,
	}

	return &out
}

// Create initializes the builder
func (app *builder) Create() Builder {
	return createBuilder()
}

// WithList adds a list to the builder
func (app *builder) WithList(list []Pointer) Builder {
	app.list = list
	return app
}

// Now builds a new Pointers instance
func (app *builder) Now() (Pointers, error) {
	if app.list != nil && len(app.list) <= 0 {
		app.list = nil
	}

	if app.list == nil {
		return nil, errors.New("there must be at least 1 Pointer in order to build a Pointers instance")
	}

	// order by index
	lastIndex := uint64(0)
	orderedPointers := map[uint64]Pointer{}
	for idx, onePointer := range app.list {
		delimiter := onePointer.Storage().Delimiter()
		index := delimiter.Index()
		if idx <= 0 {
			lastIndex = index
			orderedPointers[lastIndex] = onePointer
			continue
		}

		orderedPointers[index] = onePointer
		lastIndex = index
	}

	output := []Pointer{}
	for _, onePointer := range orderedPointers {
		output = append(output, onePointer)
	}

	return createPointers(
		output,
	), nil
}
