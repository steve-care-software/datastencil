package outputs

import (
	"reflect"
	"strings"
	"testing"

	"github.com/steve-care-software/webx/engine/vms/domain/instances/layers/outputs/kinds"
)

func TestOutput_Success(t *testing.T) {
	variable := "myVariable"
	kind := kinds.NewKindWithPromptForTests()
	output := NewOutputForTests(variable, kind)

	retVariable := output.Variable()
	if variable != retVariable {
		t.Errorf("the returned variable was expected to be '%s', '%s' returned", variable, retVariable)
		return
	}

	retKind := output.Kind()
	if !reflect.DeepEqual(kind, retKind) {
		t.Errorf("the returned kind is invalid")
		return
	}

	if output.HasExecute() {
		t.Errorf("the output was expected to NOT contain an execute")
		return
	}
}

func TestOutput_withExecute_Success(t *testing.T) {
	variable := "myVariable"
	kind := kinds.NewKindWithPromptForTests()
	execute := []string{
		"this is a command to execute",
	}

	output := NewOutputWithExecuteForTests(variable, kind, execute)

	retVariable := output.Variable()
	if variable != retVariable {
		t.Errorf("the returned variable was expected to be '%s', '%s' returned", variable, retVariable)
		return
	}

	retKind := output.Kind()
	if !reflect.DeepEqual(kind, retKind) {
		t.Errorf("the returned kind is invalid")
		return
	}

	if !output.HasExecute() {
		t.Errorf("the output was expected to contain an execute")
		return
	}

	retExecute := output.Execute()
	if !reflect.DeepEqual(execute, retExecute) {
		t.Errorf("the returned execute was expected to be '%s', '%s' returned", strings.Join(execute, ","), strings.Join(retExecute, ","))
		return
	}
}

func TestOutput_withoutVariable_returnsError(t *testing.T) {
	kind := kinds.NewKindWithPromptForTests()
	_, err := NewBuilder().Create().WithKind(kind).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestOutput_withoutKind_returnsError(t *testing.T) {
	variable := "myVariable"
	_, err := NewBuilder().Create().WithVariable(variable).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
