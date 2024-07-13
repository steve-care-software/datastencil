package applications

import (
	"github.com/steve-care-software/datastencil/domain/instances/executions"
	"github.com/steve-care-software/historydb/domain/hash"
)

const invalidPatternErr = "the provided context (%d) does not exists"

// Application represents an application
type Application interface {
	List() ([][]string, error)
	Init(dbPath []string, name string, description string) (*uint, error)
	Begin(dbPath []string) (*uint, error)
	Execute(context uint, input []byte) ([]byte, error)
	ExecuteWithPath(context uint, inputPath []string) ([]byte, error)
	ExecuteLayer(context uint, input []byte, layerPath []string) ([]byte, error)
	ExecuteLayerWithPath(context uint, inputPath []string, layerPath []string) ([]byte, error)
	RetrieveAll(context uint, index uint, length uint) (executions.Executions, error)
	RetrieveAt(context uint, index uint) (executions.Execution, error)
	Amount(context uint) (*uint, error)
	Head(context uint) (hash.Hash, error)
	Commit(context uint) error
	Rollback(context uint) error
	Cancel(context uint) error
	Merge(baseContext uint, content uint) error
}
