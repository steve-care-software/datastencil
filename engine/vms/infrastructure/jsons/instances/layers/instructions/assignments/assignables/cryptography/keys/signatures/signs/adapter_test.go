package signs

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
)

func TestAdapter_withCreate_Success(t *testing.T) {
	ins := signs.NewSignWithCreateForTests(
		creates.NewCreateForTests(
			"myMessage",
			"myPK",
		),
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

func TestAdapter_withValidate_Success(t *testing.T) {
	ins := signs.NewSignWithValidateForTests(
		validates.NewValidateForTests(
			"mySig",
			"myMessage",
			"myPubKey",
		),
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
