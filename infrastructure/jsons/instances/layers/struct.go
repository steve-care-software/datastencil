package layers

import (
	json_instructions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/instructions"
	json_output "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/layers/outputs"
)

// Layer represents the layer
type Layer struct {
	Instructions []json_instructions.Instruction `json:"instructions"`
	Output       json_output.Output              `json:"output"`
	Input        string                          `json:"input"`
}
