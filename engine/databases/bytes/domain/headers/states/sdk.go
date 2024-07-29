package states

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/headers/states/containers"
)

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewStateBuilder creates a new state builder
func NewStateBuilder() StateBuilder {
	return createStateBuilder()
}

// Adapter represents an state adapter
type Adapter interface {
	InstanceToBytes(ins State) ([]byte, error)
	BytesToInstance(data []byte) (State, error)
	InstancesToBytes(ins States) ([]byte, error)
	BytesToInstances(data []byte) (States, error)
}

// Builder represents a states builder
type Builder interface {
	Create() Builder
	WithList(list []State) Builder
	Now() (States, error)
}

// States represents a states
type States interface {
	List() []State
}

// StateBuilder represents a state builder
type StateBuilder interface {
	Create() StateBuilder
	WithContainers(containers containers.Containers) StateBuilder
	IsDeleted() StateBuilder
	Now() (State, error)
}

// State represents an state
type State interface {
	IsDeleted() bool
	HasContainers() bool
	Containers() containers.Containers
}
