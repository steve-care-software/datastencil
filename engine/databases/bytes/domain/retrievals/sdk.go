package retrievals

// NewBuilder creates a new builder
func NewBuilder() Builder {
	return createBuilder()
}

// NewRetrievalBuilder creates a new retrieval builder
func NewRetrievalBuilder() RetrievalBuilder {
	return createRetrievalBuilder()
}

// Adapter represents a retrieval adapter
type Adapter interface {
	InstancesToBytes(ins Retrievals) ([]byte, error)
	BytesToInstances(data []byte) (Retrievals, []byte, error)
	InstanceToBytes(ins Retrieval) ([]byte, error)
	BytesToInstance(data []byte) (Retrieval, []byte, error)
}

// Builder represents the retrievals builder
type Builder interface {
	Create() Builder
	WithList(list []Retrieval) Builder
	Now() (Retrievals, error)
}

// Retrievals represents retrievals
type Retrievals interface {
	List() []Retrieval
}

// RetrievalBuilder represents the retrieval builder
type RetrievalBuilder interface {
	Create() RetrievalBuilder
	WithIndex(index uint64) RetrievalBuilder
	WithLength(length uint64) RetrievalBuilder
	Now() (Retrieval, error)
}

// Retrieval represents a retrieval
type Retrieval interface {
	Index() uint64
	Length() uint64
}
