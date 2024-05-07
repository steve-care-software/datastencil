package instances

import (
	"github.com/steve-care-software/datastencil/domain/hash"
)

// Adapter represents the instance adapter
type Adapter interface {
	ToBytes(ins Instance) ([]byte, error)
	ToInstance(data []byte) (Instance, error)
}

// Instance represents an instance
type Instance interface {
	Hash() hash.Hash
}

// Repository represents an instance repository
type Repository interface {
	Dir(path []string) ([]string, error)
	List(path []string) ([]string, error)
	Exists(path []string) (Instance, error)
	Retrieve(path []string) (Instance, error)
}

// Service represents an instance service
type Service interface {
	Begin() (*uint, error)
	Insert(context uint, path []string, instance Instance) error
	Delete(context uint, path []string) error
	Commit(context uint) error
	Cancel(context uint) error
}
