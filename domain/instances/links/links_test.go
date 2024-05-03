package links

import (
	"testing"

	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

func TestLinks_withList_Success(t *testing.T) {
	pFirstLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for first layer"))
	pSecondLayer, _ := hash.NewAdapter().FromBytes([]byte("this is some bytes for second layer"))
	list := []Link{
		NewLinkForTests(
			origins.NewOriginForTests(
				resources.NewResourceForTests(*pFirstLayer),
				operators.NewOperatorWithAndForTests(),
				origins.NewValueWithResourceForTests(
					resources.NewResourceForTests(*pSecondLayer),
				),
			),
			elements.NewElementsForTests([]elements.Element{
				elements.NewElementForTests(
					layers.NewLayerForTests(
						instructions.NewInstructionsForTests([]instructions.Instruction{
							instructions.NewInstructionWithStopForTests(),
						}),
						outputs.NewOutputForTests(
							"myVariable",
							kinds.NewKindWithContinueForTests(),
						),
						"myInput",
					),
				),
			}),
		),
	}

	ins := NewLinksForTests(list)
	retList := ins.List()
	if len(list) != len(retList) {
		t.Errorf("the returned list is invalid")
		return
	}
}

func TestLinks_withEmptyList_returnsError(t *testing.T) {
	list := []Link{}
	_, err := NewBuilder().Create().WithList(list).Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}

func TestLinks_withoutList_returnsError(t *testing.T) {
	_, err := NewBuilder().Create().Now()
	if err == nil {
		t.Errorf("the error was expected to be valid")
		return
	}
}