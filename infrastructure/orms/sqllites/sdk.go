package sqllites

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"

	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons/connections"
	"github.com/steve-care-software/datastencil/domain/instances/commits/actions/resources/instances/skeletons/resources"
	"github.com/steve-care-software/datastencil/domain/hash"
	"github.com/steve-care-software/datastencil/domain/instances/libraries"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/bytes"
	layers_bytes "github.com/steve-care-software/datastencil/domain/instances/links/layers/instructions/assignments/assignables/bytes"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs"
	"github.com/steve-care-software/datastencil/domain/instances/links/layers/outputs/kinds"
	"github.com/steve-care-software/datastencil/domain/instances/links"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements"
	"github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions"
	conditions_resources "github.com/steve-care-software/datastencil/domain/instances/links/elements/conditions/resources"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins"
	"github.com/steve-care-software/datastencil/domain/instances/links/origins/operators"
	links_resources "github.com/steve-care-software/datastencil/domain/instances/links/origins/resources"
)

// listInstanceToElementHashesListFn takes a list instance and return its element's hashes
type listInstanceToElementHashesListFn func(ins instances.Instance) ([]hash.Hash, error)

// elementsToListInstanceFn takes a list of elements and returns a list instance
type elementsToListInstanceFn func(input []interface{}) (instances.Instance, error)

// buildInstanceFn takes values and build an Instance instance
type buildInstanceFn func(values map[string]interface{}) (instances.Instance, error)

// callMethodOnInstanceFn calls a method related to a field on instance, then returns its value
type callMethodOnInstanceFn func(ins instances.Instance, fieldName string) (bool, interface{}, error)

const resourceNameDelimiter = "_"
const endOfLine = "\n"
const connectionNameDelimiter = "$"

// NewOrmRepository creates a new orm repository
func NewOrmRepository(
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
) instances.RepositoryBackup {
	hashAdapter := hash.NewAdapter()
	buildInstances := map[string]buildInstanceFn{
		"library": func(values map[string]interface{}) (instances.Instance, error) {
			builder := libraries.NewBuilder()
			if value, ok := values["layers"]; ok {
				if value, ok := value.(layers.Layers); ok {
					builder.WithLayers(value)
				}
			}

			if value, ok := values["links"]; ok {
				if value, ok := value.(links.Links); ok {
					builder.WithLinks(value)
				}
			}

			return builder.Now()
		},
		"library_link": func(values map[string]interface{}) (instances.Instance, error) {
			builder := links.NewLinkBuilder()
			if value, ok := values["origin"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithOrigin(ins.(origins.Origin))
					}
				}
			}

			if value, ok := values["elements"]; ok {
				if value, ok := value.(elements.Elements); ok {
					builder.WithElements(value)
				}
			}

			return builder.Now()
		},
		"library_link_element": func(values map[string]interface{}) (instances.Instance, error) {
			builder := elements.NewElementBuilder()
			if value, ok := values["layer"]; ok {
				if value != nil {
					builder.WithLayerBytes(value.([]byte))
				}
			}

			if value, ok := values["condition"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithCondition(ins.(conditions.Condition))
					}
				}
			}

			return builder.Now()
		},
		"library_link_element_condition": func(values map[string]interface{}) (instances.Instance, error) {
			builder := conditions.NewBuilder()
			if value, ok := values["resource"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithResource(ins.(conditions_resources.Resource))
					}
				}
			}

			if value, ok := values["next"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithNext(ins.(conditions.ConditionValue))
					}
				}
			}

			return builder.Now()
		},
		"library_link_element_condition_value": func(values map[string]interface{}) (instances.Instance, error) {
			builder := conditions.NewConditionValueBuilder()
			if value, ok := values["resource"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithResource(ins.(conditions_resources.Resource))
					}
				}
			}

			if value, ok := values["condition"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithCondition(ins.(conditions.Condition))
					}
				}
			}

			return builder.Now()
		},
		"library_link_element_condition_resource": func(values map[string]interface{}) (instances.Instance, error) {
			builder := conditions_resources.NewBuilder()
			if value, ok := values["code"]; ok {
				if casted, ok := value.(int64); ok {
					builder.WithCode(uint(casted))
				}
			}

			if value, ok := values["is_raised_in_layer"]; ok {
				if value.(int64) != 0 {
					builder.IsRaisedInLayer()
				}
			}

			return builder.Now()
		},
		"library_link_origin": func(values map[string]interface{}) (instances.Instance, error) {
			builder := origins.NewBuilder()
			if value, ok := values["resource"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithResource(ins.(links_resources.Resource))
					}
				}
			}

			if value, ok := values["operator"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithOperator(ins.(operators.Operator))
					}
				}
			}

			if value, ok := values["next"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithNext(ins.(origins.Value))
					}
				}
			}

			return builder.Now()
		},
		"library_link_origin_value": func(values map[string]interface{}) (instances.Instance, error) {
			builder := origins.NewValueBuilder()
			if value, ok := values["resource"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithResource(ins.(links_resources.Resource))
					}
				}
			}

			if value, ok := values["origin"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithOrigin(ins.(origins.Origin))
					}
				}
			}

			return builder.Now()
		},
		"library_link_origin_resource": func(values map[string]interface{}) (instances.Instance, error) {
			builder := links_resources.NewBuilder()
			if value, ok := values["layer"]; ok {
				if casted, ok := value.([]byte); ok {
					pHash, err := hash.NewAdapter().FromBytes(casted)
					if err != nil {
						return nil, err
					}

					builder.WithLayer(*pHash)
				}
			}

			if value, ok := values["is_mandatory"]; ok {
				if value.(int64) != 0 {
					builder.IsMandatory()
				}
			}

			return builder.Now()
		},
		"library_link_origin_operator": func(values map[string]interface{}) (instances.Instance, error) {
			builder := operators.NewBuilder()
			if value, ok := values["is_and"]; ok {
				if value.(int64) != 0 {
					builder.IsAnd()
				}
			}

			if value, ok := values["is_or"]; ok {
				if value.(int64) != 0 {
					builder.IsOr()
				}
			}

			if value, ok := values["is_xor"]; ok {
				if value.(int64) != 0 {
					builder.IsXor()
				}
			}

			return builder.Now()
		},
		"library_layer": func(values map[string]interface{}) (instances.Instance, error) {
			builder := layers.NewLayerBuilder()
			if value, ok := values["instructions"]; ok {
				if ins, ok := value.(instructions.Instructions); ok {
					builder.WithInstructions(ins.(instructions.Instructions))
				}
			}

			if value, ok := values["output"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithOutput(ins.(outputs.Output))
					}
				}
			}

			if value, ok := values["input"]; ok {
				if value != nil {
					builder.WithInput(value.(string))
				}
			}

			return builder.Now()
		},
		"library_layer_output": func(values map[string]interface{}) (instances.Instance, error) {
			builder := outputs.NewBuilder()
			if value, ok := values["variable"]; ok {
				builder.WithVariable(value.(string))
			}

			if value, ok := values["kind"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithKind(ins.(kinds.Kind))
					}
				}
			}

			if value, ok := values["execute"]; ok {
				if value != nil {
					builder.WithExecute(value.(string))
				}
			}

			return builder.Now()
		},
		"library_layer_output_kind": func(values map[string]interface{}) (instances.Instance, error) {
			builder := kinds.NewBuilder()
			if value, ok := values["prompt"]; ok {
				if value.(int64) != 0 {
					builder.IsPrompt()
				}
			}

			if value, ok := values["continue"]; ok {
				if value.(int64) != 0 {
					builder.IsContinue()
				}
			}

			return builder.Now()
		},
		"library_layer_instruction": func(values map[string]interface{}) (instances.Instance, error) {
			builder := instructions.NewInstructionBuilder()
			if value, ok := values["assignment"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithAssignment(ins.(assignments.Assignment))
					}
				}
			}

			return builder.Now()
		},
		"library_layer_instruction_assignment": func(values map[string]interface{}) (instances.Instance, error) {
			builder := assignments.NewBuilder()
			if value, ok := values["name"]; ok {
				builder.WithName(value.(string))
			}

			if value, ok := values["assignable"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithAssignable(ins.(assignables.Assignable))
					}
				}
			}

			return builder.Now()
		},
		"library_layer_instruction_assignment_assignable": func(values map[string]interface{}) (instances.Instance, error) {
			builder := assignables.NewBuilder()
			if value, ok := values["bytes"]; ok {
				if pIns, ok := value.(*instances.Instance); ok {
					if pIns != nil {
						ins := *pIns
						builder.WithBytes(ins.(bytes.Bytes))
					}
				}
			}

			return builder.Now()
		},
		"library_layer_instruction_assignment_assignable_bytes": func(values map[string]interface{}) (instances.Instance, error) {
			builder := layers_bytes.NewBuilder()
			if value, ok := values["joins"]; ok {
				if value != nil {
					split := strings.Split(value.(string), resourceNameDelimiter)
					builder.WithJoin(split)
				}

			}

			if value, ok := values["compares"]; ok {
				if value != nil {
					split := strings.Split(value.(string), resourceNameDelimiter)
					builder.WithCompare(split)
				}

			}

			if value, ok := values["hash_bytes"]; ok {
				if value != nil {
					builder.WithHashBytes(value.(string))
				}
			}

			return builder.Now()
		},
	}

	listInstances := map[string]elementsToListInstanceFn{
		"library_links": func(input []interface{}) (instances.Instance, error) {
			output := []links.Link{}
			for _, oneIns := range input {
				output = append(output, oneIns.(links.Link))
			}

			return links.NewBuilder().Create().
				WithList(output).
				Now()
		},
		"link_elements": func(input []interface{}) (instances.Instance, error) {
			output := []elements.Element{}
			for _, oneIns := range input {
				output = append(output, oneIns.(elements.Element))
			}

			return elements.NewBuilder().Create().
				WithList(output).
				Now()
		},
		"library_layers": func(input []interface{}) (instances.Instance, error) {
			output := []layers.Layer{}
			for _, oneIns := range input {
				output = append(output, oneIns.(layers.Layer))
			}

			return layers.NewBuilder().Create().
				WithList(output).
				Now()
		},
		"layer_instructions": func(input []interface{}) (instances.Instance, error) {
			output := []instructions.Instruction{}
			for _, oneIns := range input {
				output = append(output, oneIns.(instructions.Instruction))
			}

			return instructions.NewBuilder().Create().
				WithList(output).
				Now()
		},
	}

	return createOrmRepository(
		hashAdapter,
		buildInstances,
		listInstances,
		skeleton,
		dbPtr,
	)
}

// NewOrmService creates a new orm service
func NewOrmService(
	repository instances.RepositoryBackup,
	skeleton skeletons.Skeleton,
	dbPtr *sql.DB,
	txPtr *sql.Tx,
) instances.Service {

	callMethodsOnInstances := map[string]callMethodOnInstanceFn{
		"library": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(libraries.Library); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "layers":
					return true, casted.Layers(), nil
				case "links":
					if casted.HasLinks() {
						return true, casted.Links(), nil
					}

					return false, nil, nil

				}

				str := fmt.Sprintf("link: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a Library instance")
		},
		"library_link": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(links.Link); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "origin":
					return true, casted.Origin().Hash().Bytes(), nil
				case "elements":
					return true, casted.Elements(), nil
				}

				str := fmt.Sprintf("link: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a Link instance")
		},
		"library_link_element": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(elements.Element); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "layer":
					return true, casted.Layer().Bytes(), nil
				case "condition":
					if casted.HasCondition() {
						return true, casted.Condition().Hash().Bytes(), nil
					}

					return false, nil, nil
				}

				str := fmt.Sprintf("element: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Element instance")
		},
		"library_link_element_condition": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(conditions.Condition); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "resource":
					return true, casted.Resource().Hash().Bytes(), nil
				case "next":
					if casted.HasNext() {
						return true, casted.Next().Hash().Bytes(), nil
					}

					return false, nil, nil
				}

				str := fmt.Sprintf("condition: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a Condition instance")
		},
		"library_link_element_condition_value": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(conditions.ConditionValue); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "resource":
					if casted.IsResource() {
						return true, casted.Resource().Hash().Bytes(), nil
					}

					return false, nil, nil
				case "condition":
					if casted.IsCondition() {
						return true, casted.Condition().Hash().Bytes(), nil
					}

					return false, nil, nil
				}

				str := fmt.Sprintf("conditionValue: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a ConditionValue instance")
		},
		"library_link_element_condition_resource": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(conditions_resources.Resource); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "code":
					return true, casted.Code(), nil
				case "is_raised_in_layer":
					return true, casted.IsRaisedInLayer(), nil
				}

				str := fmt.Sprintf("conditionresource: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a ConditionResource instance")
		},
		"library_link_origin": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(origins.Origin); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "resource":
					return true, casted.Resource().Hash().Bytes(), nil
				case "operator":
					return true, casted.Operator().Hash().Bytes(), nil
				case "next":
					return true, casted.Next().Hash().Bytes(), nil
				}

				str := fmt.Sprintf("origin: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Origin instance")
		},
		"library_link_origin_value": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(origins.Value); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "resource":
					if casted.IsResource() {
						return true, casted.Resource().Hash().Bytes(), nil
					}

					return false, nil, nil
				case "origin":
					if casted.IsOrigin() {
						return true, casted.Origin().Hash().Bytes(), nil
					}

					return false, nil, nil
				}

				str := fmt.Sprintf("resource: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Value instance")
		},
		"library_link_origin_resource": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(links_resources.Resource); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "layer":
					return true, casted.Layer().Bytes(), nil
				case "is_mandatory":
					return true, casted.IsMandatory(), nil
				}

				str := fmt.Sprintf("resource: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Resource instance")
		},
		"library_link_origin_operator": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(operators.Operator); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "is_and":
					return true, casted.IsAnd(), nil
				case "is_or":
					return true, casted.IsOr(), nil
				case "is_xor":
					return true, casted.IsXor(), nil
				}

				str := fmt.Sprintf("operator: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Operator instance")
		},
		"library_layer": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(layers.Layer); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "instructions":
					return true, casted.Instructions(), nil
				case "output":
					return true, casted.Output().Hash().Bytes(), nil
				case "input":
					return true, casted.Input(), nil
				}

				str := fmt.Sprintf("layer: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a Layer instance")
		},
		"library_layer_output": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(outputs.Output); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "variable":
					return true, casted.Variable(), nil
				case "kind":
					return true, casted.Kind().Hash().Bytes(), nil
				case "execute":
					return true, casted.Execute(), nil
				}

				str := fmt.Sprintf("output: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Output instance")
		},
		"library_layer_output_kind": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(kinds.Kind); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "prompt":
					return true, casted.IsPrompt(), nil
				case "continue":
					return true, casted.IsContinue(), nil
				}

				str := fmt.Sprintf("kind: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a Kind instance")
		},
		"library_layer_instruction": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(instructions.Instruction); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "assignment":
					return true, casted.Assignment().Hash().Bytes(), nil
				}

				str := fmt.Sprintf("instruction: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Instruction instance")
		},
		"library_layer_instruction_assignment": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(assignments.Assignment); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "name":
					return true, casted.Name(), nil
				case "assignable":
					return true, casted.Assignable().Hash().Bytes(), nil
				}

				str := fmt.Sprintf("assignment: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Assignment instance")
		},
		"library_layer_instruction_assignment_assignable": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(assignables.Assignable); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "bytes":
					return casted.IsBytes(), casted.Bytes().Hash().Bytes(), nil
				}

				str := fmt.Sprintf("assignable: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain an Assignable instance")
		},
		"library_layer_instruction_assignment_assignable_bytes": func(ins instances.Instance, fieldName string) (bool, interface{}, error) {
			if casted, ok := ins.(bytes.Bytes); ok {
				switch fieldName {
				case "hash":
					return true, casted.Hash().Bytes(), nil
				case "joins":
					return casted.IsJoin(), strings.Join(casted.Join(), resourceNameDelimiter), nil
				case "compares":
					return casted.IsCompare(), strings.Join(casted.Compare(), resourceNameDelimiter), nil
				case "hash_bytes":
					return casted.IsHashBytes(), casted.HashBytes(), nil
				}

				str := fmt.Sprintf("bytes: the fieldName is invalid: %s", fieldName)
				return false, nil, errors.New(str)
			}

			return false, nil, errors.New("the Instance was expected to contain a Bytes instance")
		},
	}

	listInstanceToElementHashesListFn := map[string]listInstanceToElementHashesListFn{
		"library_layers": func(ins instances.Instance) ([]hash.Hash, error) {
			if ins, ok := ins.(layers.Layers); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the Instance was expected to contain a Layers instance")
		},
		"layer_instructions": func(ins instances.Instance) ([]hash.Hash, error) {
			if ins, ok := ins.(instructions.Instructions); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the Instance was expected to contain an Instructions instance")
		},
		"library_links": func(ins instances.Instance) ([]hash.Hash, error) {
			if ins, ok := ins.(links.Links); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the Instance was expected to contain a Links instance")
		},
		"link_elements": func(ins instances.Instance) ([]hash.Hash, error) {
			if ins, ok := ins.(elements.Elements); ok {
				output := []hash.Hash{}
				list := ins.List()
				for _, oneIns := range list {
					output = append(output, oneIns.Hash())
				}

				return output, nil
			}

			return nil, errors.New("the Instance was expected to contain an Elements instance")
		},
	}

	hashAdapter := hash.NewAdapter()
	return createOrmService(
		callMethodsOnInstances,
		listInstanceToElementHashesListFn,
		repository,
		hashAdapter,
		skeleton,
		dbPtr,
		txPtr,
	)
}

// NewSkeletonFactory creates a new skeleton factory
func NewSkeletonFactory() skeletons.Factory {
	builder := skeletons.NewBuilder()
	resourcesBuilder := resources.NewBuilder()
	resourceBuilder := resources.NewResourceBuilder()
	fieldsBuilder := resources.NewFieldsBuilder()
	fieldBuilder := resources.NewFieldBuilder()
	kindBuilder := resources.NewKindBuilder()
	nativeBuilder := resources.NewNativeBuilder()
	listBuilder := resources.NewListBuilder()
	connectionsBuilder := connections.NewBuilder()
	connectionBuilder := connections.NewConnectionBuilder()
	connectionFieldBuilder := connections.NewFieldBuilder()
	return createSkeletonFactory(
		builder,
		resourcesBuilder,
		resourceBuilder,
		fieldsBuilder,
		fieldBuilder,
		kindBuilder,
		nativeBuilder,
		listBuilder,
		connectionsBuilder,
		connectionBuilder,
		connectionFieldBuilder,
	)
}
