package commands

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands/results"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
)

type command struct {
	hash   hash.Hash
	input  []byte
	layer  layers.Layer
	result results.Result
	parent Link
}

func createCommand(
	hash hash.Hash,
	input []byte,
	layer layers.Layer,
	result results.Result,
	parent Link,
) Command {
	return createCommandInternally(
		hash,
		input,
		layer,
		result,
		parent,
	)
}

func createCommandInternally(
	hash hash.Hash,
	input []byte,
	layer layers.Layer,
	result results.Result,
	parent Link,
) Command {
	out := command{
		hash:   hash,
		input:  input,
		layer:  layer,
		result: result,
		parent: parent,
	}

	return &out
}

// Hash returns the hash
func (obj *command) Hash() hash.Hash {
	return obj.hash
}

// Input returns the input
func (obj *command) Input() []byte {
	return obj.input
}

// Layer returns the layer
func (obj *command) Layer() layers.Layer {
	return obj.layer
}

// Result returns the result
func (obj *command) Result() results.Result {
	return obj.result
}

// Parent returns the parent, if any
func (obj *command) Parent() Link {
	return obj.parent
}
