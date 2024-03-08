package layers

import (
	"reflect"
	"testing"

	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/libraries/layers/instructions/assignments/assignables/bytes"
)

func TestAssignment_Success(t *testing.T) {
	name := "myName"
	assignable := assignables.NewAssignableWithBytesForTests(bytes.NewBytesWithJoinForTests([]string{
		"first",
		"second",
	}))

	assignment := NewAssignmentForTests(name, assignable)
	retName := assignment.Name()
	if name != retName {
		t.Errorf("the name was expected to be '%s', '%s' returned", name, retName)
		return
	}

	retAssignable := assignment.Assignable()
	if !reflect.DeepEqual(assignable, retAssignable) {
		t.Errorf("the assignable is invalid")
		return
	}
}

func TestAssignment_withoutName_returnsError(t *testing.T) {
	assignable := assignables.NewAssignableWithBytesForTests(bytes.NewBytesWithJoinForTests([]string{
		"first",
		"second",
	}))

	_, err := NewAssignmentBuilder().Create().WithAssignable(assignable).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}

func TestAssignment_withoutAssignable_returnsError(t *testing.T) {
	name := "myName"
	_, err := NewAssignmentBuilder().Create().WithName(name).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid, nil returned")
		return
	}
}
