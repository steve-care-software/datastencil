package conditions

import (
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions/resources"
)

type condition struct {
	hash     hash.Hash
	resource resources.Resource
	next     Condition
}

func createCondition(
	hash hash.Hash,
	resource resources.Resource,
) Condition {
	return createConditionInternally(hash, resource, nil)
}

func createConditionWithNext(
	hash hash.Hash,
	resource resources.Resource,
	next Condition,
) Condition {
	return createConditionInternally(hash, resource, next)
}

func createConditionInternally(
	hash hash.Hash,
	resource resources.Resource,
	next Condition,
) Condition {
	out := condition{
		hash:     hash,
		resource: resource,
		next:     next,
	}

	return &out
}

// Hash returns the hash
func (obj *condition) Hash() hash.Hash {
	return obj.hash
}

// Resource returns the resource
func (obj *condition) Resource() resources.Resource {
	return obj.resource
}

// HasNext returns true if there is a next, false otherwise
func (obj *condition) HasNext() bool {
	return obj.next != nil
}

// Next returns the next value
func (obj *condition) Next() Condition {
	return obj.next
}
