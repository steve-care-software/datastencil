package conditions

import (
	"path/filepath"

	"github.com/steve-care-software/datastencil/domain/hash"
)

type resource struct {
	hash         hash.Hash
	path         []string
	mustBeLoaded bool
}

func createResource(
	hash hash.Hash,
	path []string,
	mustBeLoaded bool,
) Resource {
	out := resource{
		hash:         hash,
		path:         path,
		mustBeLoaded: mustBeLoaded,
	}

	return &out
}

// Hash returns the hash
func (obj *resource) Hash() hash.Hash {
	return obj.hash
}

// Path returns the path
func (obj *resource) Path() []string {
	return obj.path
}

// MustBeLoaded returns true if must be loaded, false otherwise
func (obj *resource) MustBeLoaded() bool {
	return obj.mustBeLoaded
}

// Match returns true if there is a match, false otherwise
func (obj *resource) Match(history [][]string) bool {
	isLoaded := false
	path := filepath.Join(obj.path...)
	for _, oneHistory := range history {
		historyPath := filepath.Join(oneHistory...)
		if historyPath == path {
			isLoaded = true
			break
		}
	}

	if !isLoaded && obj.mustBeLoaded {
		return false
	}

	return true
}
