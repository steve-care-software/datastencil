package cryptography

import (
	"bytes"
	"testing"

	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography/decrypts"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/cryptography/encrypts"
)

func TestAdapter_withEncrypt_Success(t *testing.T) {
	ins := cryptography.NewCryptographyWithEncryptForTests(
		encrypts.NewEncryptForTests(
			"myMessage",
			"myPassword",
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

func TestAdapter_withDecrypt_Success(t *testing.T) {
	ins := cryptography.NewCryptographyWithDecryptForTests(
		decrypts.NewDecryptForTests(
			"myCipher",
			"myPassword",
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
