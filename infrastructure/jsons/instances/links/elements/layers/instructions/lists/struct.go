package lists

import (
	json_deletes "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/lists/deletes"
	json_inserts "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/links/elements/layers/instructions/lists/inserts"
)

// List represents a list
type List struct {
	Delete *json_deletes.Delete `json:"delete"`
	Insert *json_inserts.Insert `json:"insert"`
}
