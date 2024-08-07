package votes

import (
	json_creates "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	json_validates "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
)

// Vote represents a vote
type Vote struct {
	Create   *json_creates.Create     `json:"create"`
	Validate *json_validates.Validate `json:"validate"`
}
