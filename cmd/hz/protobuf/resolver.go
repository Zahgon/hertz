package protobuf

import (
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Symbol struct {
	Space   string
	Name    string
	IsValue bool
	Type    *model.Type
	Value   interface{}
	Scope   *descriptorpb.FileDescriptorProto
}

type NameSpace map[string]*Symbol

var (
	ConstTrue = Symbol{
		IsValue: true,
		Type:    model.TypeBool,
		Value:   true,
		Scope:   &BaseProto,
	}
	ConstFalse = Symbol{
		IsValue: true,
		Type:    model.TypeBool,
		Value:   false,
		Scope:   &BaseProto,
	}
	ConstEmptyString = Symbol{
		IsValue: true,
		Type:    model.TypeString,
		Value:   "",
		Scope:   &BaseProto,
	}
)

type PackageReference struct {
	IncludeBase string
	IncludePath string
	Model       *model.Model
	Ast         *descriptorpb.FileDescriptorProto
	Referred    bool
}

func getReferPkgMap(pkgMap map[string]string, incs []*descriptorpb.FileDescriptorProto, mainModel *model.Model) (map[string]*PackageReference, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type FileInfos struct {
	Official  map[string]*descriptorpb.FileDescriptorProto
	PbReflect map[string]*desc.FileDescriptor
}

type Resolver struct {
	rootName string
	root     NameSpace
	deps     map[string]NameSpace

	mainPkg PackageReference
	refPkgs map[string]*PackageReference

	files FileInfos
}

func updateFiles(fileName string, files FileInfos) (FileInfos, error) {
	_ = "STUB: not implemented"
	return *new(FileInfos), nil
}

func NewResolver(ast *descriptorpb.FileDescriptorProto, files FileInfos, model *model.Model, pkgMap map[string]string) (*Resolver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) GetRefModel(includeBase string) (*model.Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) getBaseType(f *descriptorpb.FieldDescriptorProto, nested []*descriptorpb.DescriptorProto) (*model.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func IsMapEntry(nt *descriptorpb.DescriptorProto) bool { _ = "STUB: not implemented"; return false }

func checkListType(typ *model.Type, label descriptorpb.FieldDescriptorProto_Label) *model.Type {
	_ = "STUB: not implemented"
	return nil
}

func getNestedType(f *descriptorpb.FieldDescriptorProto, nested []*descriptorpb.DescriptorProto) *descriptorpb.DescriptorProto {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) ResolveType(f *descriptorpb.FieldDescriptorProto, nested []*descriptorpb.DescriptorProto) (*model.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) ResolveIdentifier(id string) (ret *Symbol, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) getFieldType(f *descriptorpb.FieldDescriptorProto, nested []*descriptorpb.DescriptorProto) (*model.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (resolver *Resolver) Get(name string) *Symbol { _ = "STUB: not implemented"; return nil }

func (resolver *Resolver) ExportReferred(all, needMain bool) (ret []*PackageReference) {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) LoadAll(ast *descriptorpb.FileDescriptorProto) error {
	_ = "STUB: not implemented"
	return nil
}

func mergeNamespace(first, second NameSpace) NameSpace {
	_ = "STUB: not implemented"
	return *new(NameSpace)
}

func LoadBaseIdentifier(ast *descriptorpb.FileDescriptorProto) map[string]*Symbol {
	_ = "STUB: not implemented"
	return nil
}

func (resolver *Resolver) LoadOne(ast *descriptorpb.FileDescriptorProto) (NameSpace, error) {
	_ = "STUB: not implemented"
	return *new(NameSpace), nil
}

func (resolver *Resolver) GetFiles() FileInfos { _ = "STUB: not implemented"; return *new(FileInfos) }
