package assignables

import (
	"bytes"
	"reflect"
	"testing"

	application_bytes "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/bytes"
	application_compilers "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/compilers"
	application_constants "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/constants"
	application_cryptography "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography"
	application_decrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/decrypts"
	application_encrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/encrypts"
	application_cryptography_keys "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys"
	application_encryptions "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions"
	application_keys_decrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/decrypts"
	application_keys_encrypts "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/encryptions/encrypts"
	application_signatures "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures"
	application_signs "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs"
	application_signs_creates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/creates"
	application_signs_validates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/signs/validates"
	application_votes "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes"
	application_votes_creates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/creates"
	application_votes_validates "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/cryptography/keys/signatures/votes/validates"
	application_executables "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/excutables"
	application_lists "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/lists"
	application_fetches "github.com/steve-care-software/webx/engine/vms/applications/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/webx/engine/vms/domain/instances"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables"
	assignable_bytes "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/compilers"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/constants"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography"
	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/cryptography/decrypts"
	instructions_list "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists"
	instructions_fetches "github.com/steve-care-software/webx/engine/vms/domain/instances/layers/instructions/assignments/assignables/lists/fetches"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks"
	"github.com/steve-care-software/webx/engine/vms/domain/stacks/mocks"
)

func TestExecute_withBytes_Success(t *testing.T) {
	firstVariable := "firstVar"
	firstValue := []byte("firstValue")
	secondVariable := "secondVar"
	secondValue := []byte("secondValue")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				firstVariable,
				stacks.NewAssignableWithBytesForTests(
					firstValue,
				),
			),
			stacks.NewAssignmentForTests(
				secondVariable,
				stacks.NewAssignableWithBytesForTests(
					secondValue,
				),
			),
		}),
	)

	instruction := assignables.NewAssignableWithBytesForTests(
		assignable_bytes.NewBytesWithJoinForTests([]string{
			firstVariable,
			secondVariable,
		}),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		nil,
		application_bytes.NewApplication(),
		application_constants.NewApplication(),
		application_cryptography.NewApplication(
			application_decrypts.NewApplication(
				encryptor,
			),
			application_encrypts.NewApplication(
				encryptor,
			),
			application_cryptography_keys.NewApplication(
				application_encryptions.NewApplication(
					application_keys_decrypts.NewApplication(),
					application_keys_encrypts.NewApplication(),
					1048,
				),
				application_signatures.NewApplication(
					application_votes.NewApplication(
						application_votes_creates.NewApplication(),
						application_votes_validates.NewApplication(),
					),
					application_signs.NewApplication(
						application_signs_creates.NewApplication(),
						application_signs_validates.NewApplication(),
					),
				),
			),
		),
		application_lists.NewApplication(
			application_fetches.NewApplication(),
		),
		application_executables.NewApplication(nil, nil),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retBytes := retAssignable.Bytes()
	expected := bytes.Join([][]byte{
		firstValue,
		secondValue,
	}, []byte{})

	if !bytes.Equal(expected, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}

func TestExecute_withConstant_Success(t *testing.T) {
	instruction := assignables.NewAssignableWithConstantForTests(
		constants.NewConstantWithBoolForTests(true),
	)

	frame := stacks.NewFrameForTests()
	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		nil,
		application_bytes.NewApplication(),
		application_constants.NewApplication(),
		application_cryptography.NewApplication(
			application_decrypts.NewApplication(
				encryptor,
			),
			application_encrypts.NewApplication(
				encryptor,
			),
			application_cryptography_keys.NewApplication(
				application_encryptions.NewApplication(
					application_keys_decrypts.NewApplication(),
					application_keys_encrypts.NewApplication(),
					1048,
				),
				application_signatures.NewApplication(
					application_votes.NewApplication(
						application_votes_creates.NewApplication(),
						application_votes_validates.NewApplication(),
					),
					application_signs.NewApplication(
						application_signs_creates.NewApplication(),
						application_signs_validates.NewApplication(),
					),
				),
			),
		),
		application_lists.NewApplication(
			application_fetches.NewApplication(),
		),
		application_executables.NewApplication(nil, nil),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBool() {
		t.Errorf("the assignable was expected to contain bool")
		return
	}

	pValue := retAssignable.Bool()
	if !*pValue {
		t.Errorf("the value was expected to be true")
		return
	}
}

func TestExecute_withCryptography_Success(t *testing.T) {
	cipherVar := "cipherVar"
	cipher := []byte("this is a cipher")
	passwordVar := "passVar"
	password := []byte("this is a password")
	message := []byte("this is some message")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				cipherVar,
				stacks.NewAssignableWithBytesForTests(
					cipher,
				),
			),
			stacks.NewAssignmentForTests(
				passwordVar,
				stacks.NewAssignableWithBytesForTests(
					password,
				),
			),
		}),
	)

	instruction := assignables.NewAssignableWithCryptographyForTests(
		cryptography.NewCryptographyWithDecryptForTests(
			decrypts.NewDecryptForTests(cipherVar, passwordVar),
		),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{
			string(cipher): map[string][]byte{
				string(password): message,
			},
		},
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		nil,
		application_bytes.NewApplication(),
		application_constants.NewApplication(),
		application_cryptography.NewApplication(
			application_decrypts.NewApplication(
				encryptor,
			),
			application_encrypts.NewApplication(
				encryptor,
			),
			application_cryptography_keys.NewApplication(
				application_encryptions.NewApplication(
					application_keys_decrypts.NewApplication(),
					application_keys_encrypts.NewApplication(),
					1048,
				),
				application_signatures.NewApplication(
					application_votes.NewApplication(
						application_votes_creates.NewApplication(),
						application_votes_validates.NewApplication(),
					),
					application_signs.NewApplication(
						application_signs_creates.NewApplication(),
						application_signs_validates.NewApplication(),
					),
				),
			),
		),
		application_lists.NewApplication(
			application_fetches.NewApplication(),
		),
		application_executables.NewApplication(nil, nil),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retMessage := retAssignable.Bytes()
	if !bytes.Equal(message, retMessage) {
		t.Errorf("the returned message is invalid")
		return
	}
}

func TestExecute_WithCompiler_Success(t *testing.T) {
	compileVar := "compileVar"
	compile := []byte("this is some data")
	compiledInstance := compilers.NewCompilerWithDecompileForTests("decompileVar")

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				compileVar,
				stacks.NewAssignableWithBytesForTests(
					compile,
				),
			),
		}),
	)

	instruction := assignables.NewAssignableWithCompilerForTests(
		compilers.NewCompilerWithCompileForTests(compileVar),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{
					string(compile): compiledInstance,
				},
			),
		),
		nil,
		application_bytes.NewApplication(),
		application_constants.NewApplication(),
		application_cryptography.NewApplication(
			application_decrypts.NewApplication(
				encryptor,
			),
			application_encrypts.NewApplication(
				encryptor,
			),
			application_cryptography_keys.NewApplication(
				application_encryptions.NewApplication(
					application_keys_decrypts.NewApplication(),
					application_keys_encrypts.NewApplication(),
					1048,
				),
				application_signatures.NewApplication(
					application_votes.NewApplication(
						application_votes_creates.NewApplication(),
						application_votes_validates.NewApplication(),
					),
					application_signs.NewApplication(
						application_signs_creates.NewApplication(),
						application_signs_validates.NewApplication(),
					),
				),
			),
		),
		application_lists.NewApplication(
			application_fetches.NewApplication(),
		),
		application_executables.NewApplication(nil, nil),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsInstance() {
		t.Errorf("the assignable was expected to contain an instance")
		return
	}

	retInstance := retAssignable.Instance()
	if !reflect.DeepEqual(compiledInstance.(instances.Instance), retInstance) {
		t.Errorf("the returned instance is invalid")
		return
	}
}

func TestExecute_withList_Success(t *testing.T) {
	index := uint(2)
	value := []byte("third")

	list := stacks.NewAssignablesForTests([]stacks.Assignable{
		stacks.NewAssignableWithBytesForTests([]byte("first")),
		stacks.NewAssignableWithBytesForTests([]byte("second")),
		stacks.NewAssignableWithBytesForTests(value),
	})

	listVar := "myList"
	indexVar := "myIndex"

	frame := stacks.NewFrameWithAssignmentsForTests(
		stacks.NewAssignmentsForTests([]stacks.Assignment{
			stacks.NewAssignmentForTests(
				listVar,
				stacks.NewAssignableWithListForTests(
					list,
				),
			),
			stacks.NewAssignmentForTests(
				indexVar,
				stacks.NewAssignableWithUnsignedIntForTests(
					index,
				),
			),
		}),
	)

	instruction := assignables.NewAssignableWithListForTests(
		instructions_list.NewListWithFetchForTests(
			instructions_fetches.NewFetchForTests(listVar, indexVar),
		),
	)

	encryptor := mocks.NewEncryptor(
		map[string]map[string][]byte{},
		map[string]map[string][]byte{},
	)

	application := NewApplication(
		application_compilers.NewApplication(
			mocks.NewInstanceAdapter(
				map[string][]byte{},
				map[string]instances.Instance{},
			),
		),
		nil,
		application_bytes.NewApplication(),
		application_constants.NewApplication(),
		application_cryptography.NewApplication(
			application_decrypts.NewApplication(
				encryptor,
			),
			application_encrypts.NewApplication(
				encryptor,
			),
			application_cryptography_keys.NewApplication(
				application_encryptions.NewApplication(
					application_keys_decrypts.NewApplication(),
					application_keys_encrypts.NewApplication(),
					1048,
				),
				application_signatures.NewApplication(
					application_votes.NewApplication(
						application_votes_creates.NewApplication(),
						application_votes_validates.NewApplication(),
					),
					application_signs.NewApplication(
						application_signs_creates.NewApplication(),
						application_signs_validates.NewApplication(),
					),
				),
			),
		),
		application_lists.NewApplication(
			application_fetches.NewApplication(),
		),
		application_executables.NewApplication(nil, nil),
	)

	retAssignable, pCode, err := application.Execute(frame, instruction)
	if err != nil {
		t.Errorf("the error was expected to be nil, error returned: %s", err.Error())
		return
	}

	if pCode != nil {
		t.Errorf("the code was expected to be nil, code returned: %d", *pCode)
		return
	}

	if !retAssignable.IsBytes() {
		t.Errorf("the assignable was expected to contain bytes")
		return
	}

	retBytes := retAssignable.Bytes()
	if !bytes.Equal(value, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}
