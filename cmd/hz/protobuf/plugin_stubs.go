package protobuf

import (
	_ "unsafe"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/reflect/protoreflect"
)

const (
	FileDescriptorProto_Name_field_number             protoreflect.FieldNumber = 1
	FileDescriptorProto_Package_field_number          protoreflect.FieldNumber = 2
	FileDescriptorProto_Dependency_field_number       protoreflect.FieldNumber = 3
	FileDescriptorProto_PublicDependency_field_number protoreflect.FieldNumber = 10
	FileDescriptorProto_WeakDependency_field_number   protoreflect.FieldNumber = 11
	FileDescriptorProto_MessageType_field_number      protoreflect.FieldNumber = 4
	FileDescriptorProto_EnumType_field_number         protoreflect.FieldNumber = 5
	FileDescriptorProto_Service_field_number          protoreflect.FieldNumber = 6
	FileDescriptorProto_Extension_field_number        protoreflect.FieldNumber = 7
	FileDescriptorProto_Options_field_number          protoreflect.FieldNumber = 8
	FileDescriptorProto_SourceCodeInfo_field_number   protoreflect.FieldNumber = 9
	FileDescriptorProto_Syntax_field_number           protoreflect.FieldNumber = 12
)

const WeakFieldPrefix_goname = "XXX_weak_"

type fileInfo struct {
	*protogen.File

	allEnums      []*enumInfo
	allMessages   []*messageInfo
	allExtensions []*extensionInfo

	allEnumsByPtr         map[*enumInfo]int
	allMessagesByPtr      map[*messageInfo]int
	allMessageFieldsByPtr map[*messageInfo]*structFields

	needRawDesc bool
}

type enumInfo struct {
	*protogen.Enum

	genJSONMethod    bool
	genRawDescMethod bool
}

type messageInfo struct {
	*protogen.Message

	genRawDescMethod  bool
	genExtRangeMethod bool

	isTracked bool
	hasWeak   bool
}

type extensionInfo struct {
	*protogen.Extension
}

type structFields struct {
	count      int
	unexported map[int]string
}

func (sf *structFields) append(name string) { _ = "STUB: not implemented"; return }

type structTags [][2]string

func (tags structTags) String() string { _ = "STUB: not implemented"; return "" }

type goImportPath interface {
	String() string
	Ident(string) protogen.GoIdent
}

type trailingComment protogen.Comments

func (c trailingComment) String() string { _ = "STUB: not implemented"; return "" }

//go:linkname gotrackTags google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.gotrackTags
var gotrackTags structTags

var (
	protoPackage         goImportPath = protogen.GoImportPath("google.golang.org/protobuf/proto")
	protoifacePackage    goImportPath = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoiface")
	protoimplPackage     goImportPath = protogen.GoImportPath("google.golang.org/protobuf/runtime/protoimpl")
	protojsonPackage     goImportPath = protogen.GoImportPath("google.golang.org/protobuf/encoding/protojson")
	protoreflectPackage  goImportPath = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoreflect")
	protoregistryPackage goImportPath = protogen.GoImportPath("google.golang.org/protobuf/reflect/protoregistry")
)

//go:linkname newFileInfo google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.newFileInfo
func newFileInfo(file *protogen.File) *fileInfo

//go:linkname genPackageKnownComment google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genPackageKnownComment
func genPackageKnownComment(f *fileInfo) protogen.Comments

//go:linkname genStandaloneComments google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genStandaloneComments
func genStandaloneComments(g *protogen.GeneratedFile, f *fileInfo, n int32)

//go:linkname genGeneratedHeader google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genGeneratedHeader
func genGeneratedHeader(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo)

//go:linkname genImport google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genImport
func genImport(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo, imp protoreflect.FileImport)

//go:linkname genEnum google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genEnum
func genEnum(g *protogen.GeneratedFile, f *fileInfo, e *enumInfo)

//go:linkname genMessageInternalFields google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genMessageInternalFields
func genMessageInternalFields(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo, sf *structFields)

//go:linkname genExtensions google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genExtensions
func genExtensions(g *protogen.GeneratedFile, f *fileInfo)

//go:linkname genReflectFileDescriptor google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genReflectFileDescriptor
func genReflectFileDescriptor(gen *protogen.Plugin, g *protogen.GeneratedFile, f *fileInfo)

//go:linkname appendDeprecationSuffix google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.appendDeprecationSuffix
func appendDeprecationSuffix(prefix protogen.Comments, deprecated bool) protogen.Comments

//go:linkname genMessageDefaultDecls google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genMessageDefaultDecls
func genMessageDefaultDecls(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo)

//go:linkname genMessageKnownFunctions google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genMessageKnownFunctions
func genMessageKnownFunctions(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo)

//go:linkname genMessageMethods google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genMessageMethods
func genMessageMethods(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo)

//go:linkname genMessageOneofWrapperTypes google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.genMessageOneofWrapperTypes
func genMessageOneofWrapperTypes(g *protogen.GeneratedFile, f *fileInfo, m *messageInfo)

//go:linkname oneofInterfaceName google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.oneofInterfaceName
func oneofInterfaceName(oneof *protogen.Oneof) string

//go:linkname fieldGoType google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.fieldGoType
func fieldGoType(g *protogen.GeneratedFile, f *fileInfo, field *protogen.Field) (goType string, pointer bool)

//go:linkname fieldProtobufTagValue google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.fieldProtobufTagValue
func fieldProtobufTagValue(field *protogen.Field) string

//go:linkname fieldJSONTagValue google.golang.org/protobuf/cmd/protoc-gen-go/internal_gengo.fieldJSONTagValue
func fieldJSONTagValue(field *protogen.Field) string
