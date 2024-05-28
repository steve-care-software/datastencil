package results

import (
	json_interruptions "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commands/results/interruptions"
	json_success "github.com/steve-care-software/datastencil/infrastructure/jsons/instances/commands/results/success"
)

// Result represents result
type Result struct {
	Success      *json_success.Success            `json:"success"`
	Interruption *json_interruptions.Interruption `json:"interruption"`
}
