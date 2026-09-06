package model

type Kind uint

const (
	KindInvalid Kind = iota
	KindBool
	KindInt
	KindInt8
	KindInt16
	KindInt32
	KindInt64
	KindUint
	KindUint8
	KindUint16
	KindUint32
	KindUint64
	KindUintptr
	KindFloat32
	KindFloat64
	KindComplex64
	KindComplex128
	KindArray
	KindChan
	KindFunc
	KindInterface
	KindMap
	KindPtr
	KindSlice
	KindString
	KindStruct
	KindUnsafePointer
)

type Category int64

const (
	CategoryConstant  Category = 1
	CategoryBinary    Category = 8
	CategoryMap       Category = 9
	CategoryList      Category = 10
	CategorySet       Category = 11
	CategoryEnum      Category = 12
	CategoryStruct    Category = 13
	CategoryUnion     Category = 14
	CategoryException Category = 15
	CategoryTypedef   Category = 16
	CategoryService   Category = 17
)

type Model struct {
	FilePath string
	Package  string
	Imports  map[string]*Model

	PackageName string
	Typedefs    []TypeDef
	Constants   []Constant
	Variables   []Variable
	Functions   []Function
	Enums       []Enum
	Structs     []Struct
	Methods     []Method
	Oneofs      []Oneof
}

func (m Model) IsEmpty() bool { _ = "STUB: not implemented"; return false }

type Models []*Model

func (a *Models) MergeMap(b map[string]*Model) { _ = "STUB: not implemented"; return }

func (a *Models) MergeArray(b []*Model) { _ = "STUB: not implemented"; return }

type RequiredNess int

const (
	RequiredNess_Default  RequiredNess = 0
	RequiredNess_Required RequiredNess = 1
	RequiredNess_Optional RequiredNess = 2
)

type Type struct {
	Name     string
	Scope    *Model
	Kind     Kind
	Indirect bool
	Category Category
	Extra    []*Type
	HasNew   bool
}

func (rt *Type) ResolveDefaultValue() string { _ = "STUB: not implemented"; return "" }

func (rt *Type) ResolveNameForTypedef(scope *Model) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rt *Type) ResolveName(scope *Model) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (rt *Type) IsBinary() bool { _ = "STUB: not implemented"; return false }

func (rt *Type) IsBaseType() bool { _ = "STUB: not implemented"; return false }

func (rt *Type) IsSettable() bool { _ = "STUB: not implemented"; return false }

type TypeDef struct {
	Scope *Model
	Alias string
	Type  *Type
}

type Constant struct {
	Scope *Model
	Name  string
	Type  *Type
	Value Literal
}

type Literal interface {
	Expression() string
}

type Variable struct {
	Scope *Model
	Name  string
	Type  *Type
	Value Literal
}

type Function struct {
	Scope *Model
	Name  string
	Args  []Variable
	Rets  []Variable
	Code  string
}

type Method struct {
	Scope        *Model
	ReceiverName string
	ReceiverType *Type
	ByPtr        bool
	Function
}

type Enum struct {
	Scope  *Model
	Name   string
	GoType string
	Values []Constant
}

type Struct struct {
	Scope           *Model
	Name            string
	Fields          []Field
	Category        Category
	LeadingComments string
}

type Field struct {
	Scope            *Struct
	Name             string
	Type             *Type
	IsSetDefault     bool
	DefaultValue     Literal
	Required         RequiredNess
	Tags             Tags
	LeadingComments  string
	TrailingComments string
	IsPointer        bool
}

type Oneof struct {
	MessageName   string
	OneofName     string
	InterfaceName string
	Choices       []Choice
}

type Choice struct {
	MessageName string
	ChoiceName  string
	Type        *Type
}

type Tags []Tag

type Tag struct {
	Key       string
	Value     string
	IsDefault bool
}

func (ts Tags) String() string { _ = "STUB: not implemented"; return "" }

func (ts *Tags) Remove(name string) { _ = "STUB: not implemented"; return }

func (ts Tags) Len() int { _ = "STUB: not implemented"; return 0 }

func (ts Tags) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ts Tags) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (f Field) GenGoTags() string { _ = "STUB: not implemented"; return "" }
