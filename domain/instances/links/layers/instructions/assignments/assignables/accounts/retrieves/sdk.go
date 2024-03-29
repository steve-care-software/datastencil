package retrieves

import "github.com/steve-care-software/datastencil/domain/hash"

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// Builder represents a retrieve builder
type Builder interface {
	Create() Builder
	WithPassword(password string) Builder
	WithCredentials(credentials string) Builder
	Now() (Retrieve, error)
}

// Retrieve represents a retrieve
type Retrieve interface {
	Hash() hash.Hash
	Password() string
	Credentials() string
}
