package layers

import "github.com/steve-care-software/datastencil/domain/hash"

type assignable struct {
	hash      hash.Hash
	bytes     Bytes
	execution Execution
}

func createAssignableWithBytes(
	hash hash.Hash,
	bytes Bytes,
) Assignable {
	return createAssignableInternally(hash, bytes, nil)
}

func createAssignableWithexecution(
	hash hash.Hash,
	execution Execution,
) Assignable {
	return createAssignableInternally(hash, nil, execution)
}

func createAssignableInternally(
	hash hash.Hash,
	bytes Bytes,
	execution Execution,
) Assignable {
	out := assignable{
		hash:      hash,
		bytes:     bytes,
		execution: execution,
	}

	return &out
}

// Hash returns the hash
func (obj *assignable) Hash() hash.Hash {
	return obj.hash
}

// IsBytes returns true if there is bytes, false otherwise
func (obj *assignable) IsBytes() bool {
	return obj.bytes != nil
}

// Bytes returns the bytes, if any
func (obj *assignable) Bytes() Bytes {
	return obj.bytes
}

// IsExecution returns true if there is execution, false otherwise
func (obj *assignable) IsExecution() bool {
	return obj.execution != nil
}

// Execution returns the execution, if any
func (obj *assignable) Execution() Execution {
	return obj.execution
}
