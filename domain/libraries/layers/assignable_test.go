package layers

import (
	"reflect"
	"testing"
)

func TestAssignable_withBytes_Success(t *testing.T) {
	bytes := NewBytesWithJoinForTests([]string{
		"first",
		"second",
	})

	ins := NewAssignableWithBytesForTests(bytes)

	if !ins.IsBytes() {
		t.Errorf("the bytes was expected to contain a bytes")
		return
	}

	retBytes := ins.Bytes()
	if !reflect.DeepEqual(bytes, retBytes) {
		t.Errorf("the returned bytes is invalid")
		return
	}
}

func TestAssignable_withoutParam_returnsError(t *testing.T) {
	_, err := NewAssignableBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}
