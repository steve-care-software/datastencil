package files

import (
	"os"

	"github.com/steve-care-software/webx/engine/states/domain/hash"
	"github.com/steve-care-software/webx/engine/stencils/domain/contexts"
)

type contextRepository struct {
	adapter     contexts.Adapter
	hashAdapter hash.Adapter
	endPath     []string
}

func createContextRepository(
	adapter contexts.Adapter,
	hashAdapter hash.Adapter,
	endPath []string,
) contexts.Repository {
	out := contextRepository{
		adapter:     adapter,
		hashAdapter: hashAdapter,
		endPath:     endPath,
	}

	return &out
}

// Retrieve retrieves a context
func (app *contextRepository) Retrieve(dbPath []string) (contexts.Context, error) {
	filePath, err := prepareFilePath(app.hashAdapter, dbPath, app.endPath)
	if err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return app.adapter.ToInstance(bytes)
}
