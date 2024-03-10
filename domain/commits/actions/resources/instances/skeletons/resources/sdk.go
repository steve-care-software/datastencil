package resources

const (
	// NativeInteger represents the integer
	NativeInteger (uint8) = iota

	// NativeFloat represents the float
	NativeFloat

	// NativeString represents the string
	NativeString

	// NativeBytes represents the bytes
	NativeBytes
)

// NewBuilder creates a new builder instance
func NewBuilder() Builder {
	return createBuilder()
}

// NewResourceBuilder creates a new resource builder
func NewResourceBuilder() ResourceBuilder {
	return createResourceBuilder()
}

// NewFieldsBuilder creates a new fields builder
func NewFieldsBuilder() FieldsBuilder {
	return createFieldsBuilder()
}

// NewFieldBuilder creates a new field builder
func NewFieldBuilder() FieldBuilder {
	return createFieldBuilder()
}

// NewKindBuilder creates a new kind builder
func NewKindBuilder() KindBuilder {
	return createKindBuilder()
}

// NewNativeBuilder creates a new native builder
func NewNativeBuilder() NativeBuilder {
	return createNativeBuilder()
}

// NewListBuilder creates a new list builder
func NewListBuilder() ListBuilder {
	return createListBuilder()
}

// Builder represents a resources builder
type Builder interface {
	Create() Builder
	WithList(list []Resource) Builder
	Now() (Resources, error)
}

// Resources represents resources
type Resources interface {
	List() []Resource
	FetchByName(name string) (Resource, error)
	FetchByPath(path []string) (Resource, error)
}

// ResourceBuilder represents a resource builder
type ResourceBuilder interface {
	Create() ResourceBuilder
	WithName(name string) ResourceBuilder
	WithKey(key Field) ResourceBuilder
	WithFields(fields Fields) ResourceBuilder
	WithChildren(children Resources) ResourceBuilder
	Now() (Resource, error)
}

// Resource represents a schema resource
type Resource interface {
	Name() string
	Key() Field
	Fields() Fields
	HasChildren() bool
	Children() Resources
}

// FieldsBuilder represents a fields builder
type FieldsBuilder interface {
	Create() FieldsBuilder
	WithList(list []Field) FieldsBuilder
	Now() (Fields, error)
}

// Fields represents fields
type Fields interface {
	List() []Field
}

// FieldBuilder represents a field builder
type FieldBuilder interface {
	Create() FieldBuilder
	WithName(name string) FieldBuilder
	WithKind(kind Kind) FieldBuilder
	CanBeNil() FieldBuilder
	Now() (Field, error)
}

// Field represents a field
type Field interface {
	Name() string
	Kind() Kind
	CanBeNil() bool
}

// BuilderInstruction represents a builder instruction
type BuilderInstruction interface {
	Method() string
	ContainsParam() bool
}

// KindBuilder represents a kind builder
type KindBuilder interface {
	Create() KindBuilder
	WithNative(native Native) KindBuilder
	WithReference(reference []string) KindBuilder
	WithConnection(connection string) KindBuilder
	Now() (Kind, error)
}

// Kind represents a kind
type Kind interface {
	IsNative() bool
	Native() Native
	IsReference() bool
	Reference() []string
	IsConnection() bool
	Connection() string
}

// NativeBuilder represents a native builder
type NativeBuilder interface {
	Create() NativeBuilder
	WithSingle(single uint8) NativeBuilder
	WithList(list List) NativeBuilder
	Now() (Native, error)
}

// Native represents a native kind
type Native interface {
	IsSingle() bool
	Single() *uint8
	IsList() bool
	List() List
}

// ListBuilder represents a list builder
type ListBuilder interface {
	Create() ListBuilder
	WithValue(value uint8) ListBuilder
	WithDelimiter(delimiter string) ListBuilder
	Now() (List, error)
}

// List represents a list
type List interface {
	Value() uint8
	Delimiter() string
}