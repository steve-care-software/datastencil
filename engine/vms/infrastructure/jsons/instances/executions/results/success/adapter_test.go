package success

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/executions/results/success/outputs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs/kinds"
)

func TestAdapter_Success(t *testing.T) {

	ins := success.NewSuccessForTests(
		outputs.NewOutputForTests(
			[]byte("this is an input"),
		),
		kinds.NewKindWithPromptForTests(),
	)

	adapter := NewAdapter()

	retBytes, err := adapter.ToBytes(ins)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	retIns, err := adapter.ToInstance(retBytes)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if !bytes.Equal(ins.Hash().Bytes(), retIns.Hash().Bytes()) {
		t.Errorf("the returned instance is invalid")
		return
	}
}
