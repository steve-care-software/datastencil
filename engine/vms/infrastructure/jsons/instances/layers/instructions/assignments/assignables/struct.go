package assignables

import (
	json_bytes "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/bytes"
	json_compiler "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/compilers"
	json_constants "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/constants"
	json_cryptography "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/cryptography"
	json_executables "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executables"
	json_executions "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/executions"
	json_lists "github.com/steve-care-software/webx/engine/vms/infrastructure/jsons/instances/layers/instructions/assignments/assignables/lists"
)

// Assignable represents the assignable
type Assignable struct {
	Bytes        *json_bytes.Bytes               `json:"bytes"`
	Constant     *json_constants.Constant        `json:"constant"`
	Cryptography *json_cryptography.Cryptography `json:"cryptography"`
	Compiler     *json_compiler.Compiler         `json:"compiler"`
	List         *json_lists.List                `json:"list"`
	Execution    *json_executions.Execution      `json:"execution"`
	Executable   *json_executables.Executable    `json:"executable"`
	Variable     string                          `json:"variable"`
}
