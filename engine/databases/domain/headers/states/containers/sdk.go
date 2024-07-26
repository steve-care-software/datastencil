package containers

import (
	"github.com/steve-care-software/webx/engine/databases/domain/headers/states/containers/pointers"
)

// Adapter represents a containers adapter
type Adapter interface {
	InstancesToBytes(ins Containers) ([]byte, error)
	BytesToInstances(data []byte) (Containers, error)
	InstanceToBytes(ins Container) ([]byte, error)
	BytesToInstance(data []byte) (Container, error)
}

// Builder represents a containers builder
type Builder interface {
	Create() Builder
	WithList(list []Container) Builder
	Now() (Containers, error)
}

// Containers represents containers
type Containers interface {
	List() []Container
}

// ContrainerBuilder represents a container builder
type ContrainerBuilder interface {
	Create() ContrainerBuilder
	WithKeyname(keyname string) ContrainerBuilder
	WithPointers(pointers pointers.Pointer) ContrainerBuilder
	Now() (Container, error)
}

// Container represents a container
type Container interface {
	Keyname() string
	Pointers() pointers.Pointer
}