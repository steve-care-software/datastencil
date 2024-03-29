package layers

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(
		hashAdapter,
	)
}

// NewLayerBuilder creates a new layer builder instance
func NewLayerBuilder() LayerBuilder {
	hashAdapter := hash.NewAdapter()
	return createLayerBuilder(
		hashAdapter,
	)
}

// Adapter represents the layers adapter
type Adapter interface {
	ToData(ins Layers) ([]byte, error)
	ToInstance(data []byte) (Layers, error)
}

// Builder represents the layers builder
type Builder interface {
	Create() Builder
	WithList(list []Layer) Builder
	Now() (Layers, error)
}

// Layers represents layers
type Layers interface {
	Hash() hash.Hash
	List() []Layer
	Fetch(hash hash.Hash) (Layer, error)
}

// Repository represents the layers repository
type Repository interface {
	Retrieve(path []string) (Layers, error)
}

// LayerAdapter represents the layer adapter
type LayerAdapter interface {
	ToData(ins Layer) ([]byte, error)
	ToInstance(data []byte) (Layer, error)
}

// LayerBuilder represents a layer builder
type LayerBuilder interface {
	Create() LayerBuilder
	WithInstructions(instructions instructions.Instructions) LayerBuilder
	WithOutput(output outputs.Output) LayerBuilder
	WithInput(input string) LayerBuilder
	Now() (Layer, error)
}

// Layer represents a layer
type Layer interface {
	Hash() hash.Hash
	Instructions() instructions.Instructions
	Output() outputs.Output
	Input() string
}

// LayerRepository represents the layer repository
type LayerRepository interface {
	Retrieve(hash hash.Hash) (Layer, error)
}
