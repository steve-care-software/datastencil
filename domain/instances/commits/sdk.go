package commits

import (
	"github.com/steve-care-software/datastencil/domain/accounts/signers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions"
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	hashAdapter := hash.NewAdapter()
	return createBuilder(hashAdapter)
}

// NewContentBuilder creates a new content builder
func NewContentBuilder() ContentBuilder {
	hashAdapter := hash.NewAdapter()
	return createContentBuilder(hashAdapter)
}

// Builder represents the commit builder
type Builder interface {
	Create() Builder
	WithContent(content Content) Builder
	WithSignature(signature signers.Signature) Builder
	Now() (Commit, error)
}

// Commit represents a commit
type Commit interface {
	Hash() hash.Hash
	Content() Content
	Signature() signers.Signature
	Index() uint
	PublicKey() (signers.PublicKey, error)
}

// ContentBuilder represents the content builder
type ContentBuilder interface {
	Create() ContentBuilder
	WithActions(actions actions.Actions) ContentBuilder
	WithPrevious(previous Commit) ContentBuilder
	Now() (Content, error)
}

// Content represents a commit content
type Content interface {
	Hash() hash.Hash
	Actions() actions.Actions
	HasPrevious() bool
	Previous() Commit
}
