package executions

import "github.com/steve-care-software/webx/engine/cursors/domain/hash"

type execution struct {
	hash      hash.Hash
	excutable string
	content   Content
}

func createExecution(
	hash hash.Hash,
	excutable string,
	content Content,
) Execution {
	out := execution{
		hash:      hash,
		excutable: excutable,
		content:   content,
	}

	return &out
}

// Hash returns the hash
func (obj *execution) Hash() hash.Hash {
	return obj.hash
}

// Executable returns the executable
func (obj *execution) Executable() string {
	return obj.excutable
}

// Content returns the content
func (obj *execution) Content() Content {
	return obj.content
}
