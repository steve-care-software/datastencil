package applications

import (
	application_layers "github.com/steve-care-software/datastencil/applications/layers"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/commands"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links"
)

// NewApplication creates a new application
func NewApplication(
	execLayerApp application_layers.Application,
	repository links.LinkRepository,
	layerRepository layers.LayerRepository,
) Application {
	commandsBuilder := commands.NewBuilder()
	commandBuilder := commands.NewCommandBuilder()
	commandLinkBuilder := commands.NewLinkBuilder()
	hashAdapter := hash.NewAdapter()
	return createApplication(
		execLayerApp,
		repository,
		layerRepository,
		commandsBuilder,
		commandBuilder,
		commandLinkBuilder,
		hashAdapter,
	)
}

// Application represents a link application
type Application interface {
	Execute(input []byte) (commands.Commands, error)
	ExecuteWithContext(input []byte, context commands.Commands) (commands.Commands, error)
}
