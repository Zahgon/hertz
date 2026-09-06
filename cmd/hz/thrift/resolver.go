package thrift

import (
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/cloudwego/thriftgo/parser"
)

var (
	ConstTrue = Symbol{
		IsValue: true,
		Type:    model.TypeBool,
		Value:   true,
		Scope:   &BaseThrift,
	}
	ConstFalse = Symbol{
		IsValue: true,
		Type:    model.TypeBool,
		Value:   false,
		Scope:   &BaseThrift,
	}
	ConstEmptyString = Symbol{
		IsValue: true,
		Type:    model.TypeString,
		Value:   "",
		Scope:   &BaseThrift,
	}
)

type PackageReference struct {
	IncludeBase string
	IncludePath string
	Model       *model.Model
	Ast         *parser.Thrift
	Referred    bool
}

func getReferPkgMap(pkgMap map[string]string, incs []*parser.Include, mainModel *model.Model) (map[string]*PackageReference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type Symbol struct {
	IsValue bool
	Type    *model.Type
	Value   interface{}
	Scope   *parser.Thrift
}

type NameSpace map[string]*Symbol

type Resolver struct {
	root NameSpace
	deps map[string]NameSpace

	mainPkg PackageReference
	refPkgs map[string]*PackageReference
}

func NewResolver(ast *parser.Thrift, model *model.Model, pkgMap map[string]string) (*Resolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) GetRefModel(includeBase string) (*model.Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) getBaseType(typ *parser.Type) (*model.Type, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (resolver *Resolver) ResolveType(typ *parser.Type) (*model.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) ResolveConstantValue(constant *parser.ConstValue) (model.Literal, error) {
	_ = "STUB: not implemented"
	return *new(model.Literal), nil
}

func (resolver *Resolver) ResolveIdentifier(id string) (ret ResolvedSymbol, err error) {
	_ = "STUB: not implemented"
	return *new(ResolvedSymbol), nil
}

func (resolver *Resolver) ResolveTypeName(typ *parser.Type) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (resolver *Resolver) Get(name string) *Symbol { _ = "STUB: not implemented"; return nil }

func (resolver *Resolver) ExportReferred(all, needMain bool) (ret []*PackageReference) {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) LoadAll(ast *parser.Thrift) error { _ = "STUB: not implemented"; return nil }

func LoadBaseIdentifier() NameSpace { _ = "STUB: not implemented"; return *new(NameSpace) }

func (resolver *Resolver) LoadOne(ast *parser.Thrift) (NameSpace, error) {
	_ = "STUB: not implemented"
	return *new(NameSpace), nil
}

func switchConstantType(constant parser.ConstType) (*model.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newTypedefType(t *model.Type, name string) model.Type {
	_ = "STUB: not implemented"
	return *new(model.Type)
}
