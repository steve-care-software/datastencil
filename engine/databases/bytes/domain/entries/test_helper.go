package entries

import (
	"github.com/steve-care-software/webx/engine/databases/bytes/domain/states/containers/pointers/delimiters"
)

// NewEntriesForTests creates new entries for tests
func NewEntriesForTests(list []Entry) Entries {
	ins, err := NewBuilder().Create().WithList(list).Now()
	if err != nil {
		panic(err)
	}

	return ins
}

// NewEntryForTests creates a new entry for tests
func NewEntryForTests(keyname string, delimiter delimiters.Delimiter, bytes []byte) Entry {
	ins, err := NewEntryBuilder().Create().WithKeyname(keyname).WithDelimiter(delimiter).WithBytes(bytes).Now()
	if err != nil {
		panic(err)
	}

	return ins
}
