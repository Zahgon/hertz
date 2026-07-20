package protobuf

import (
	"github.com/cloudwego/hertz/cmd/hz/generator"
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/jhump/protoreflect/desc"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var BaseProto = descriptorpb.FileDescriptorProto{}

func getGoPackage(f *descriptorpb.FileDescriptorProto, pkgMap map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func switchBaseType(typ descriptorpb.FieldDescriptorProto_Type) *model.Type {
	_ = "STUB: not implemented"
	return nil
}

func astToService(ast *descriptorpb.FileDescriptorProto, resolver *Resolver, cmdType string, gen *protogen.Plugin) ([]*generator.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getCompatibleAnnotation(options proto.Message, anno, compatibleAnno *protoimpl.ExtensionInfo) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func parseAnnotationToClient(clientMethod *generator.ClientMethod, gen *protogen.Plugin, ast *descriptorpb.FileDescriptorProto, s *descriptorpb.ServiceDescriptorProto, m *descriptorpb.MethodDescriptorProto) error {
	_ = "STUB: not implemented"
	return nil
}

func getMethod(file *protogen.File, s *descriptorpb.ServiceDescriptorProto, m *descriptorpb.MethodDescriptorProto) (*protogen.Method, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func astToModel(ast *descriptorpb.FileDescriptorProto, rs *Resolver) (*model.Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getMessageLeadingComments(stMessage *desc.MessageDescriptor) string {
	_ = "STUB: not implemented"
	return ""
}

func getFiledComments(f *descriptorpb.FieldDescriptorProto, stMessage *desc.MessageDescriptor) (string, string) {
	_ = "STUB: not implemented"
	return "", ""
}

func formatComments(comments string) string { _ = "STUB: not implemented"; return "" }

func getNestedMessageInfoMap(stMessage *desc.MessageDescriptor) map[string]*desc.MessageDescriptor {
	_ = "STUB: not implemented"
	return nil
}

func parseDefaultValue(typ descriptorpb.FieldDescriptorProto_Type, val string) (model.Literal, error) {
	_ = "STUB: not implemented"
	return *new(model.Literal), nil
}

func isPointer(f *descriptorpb.FieldDescriptorProto, isProto3 bool) bool {
	_ = "STUB: not implemented"
	return false
}

func resolveOneof(stMessage *desc.MessageDescriptor, oneofMap map[string]model.Field, rs *Resolver, isProto3 bool, s model.Struct, ns []*descriptorpb.DescriptorProto) ([]model.Oneof, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getNewFieldName(fieldName string, fieldNameSet map[string]bool) string {
	_ = "STUB: not implemented"
	return ""
}

func checkDuplicatedFileName(vs []model.Field) { _ = "STUB: not implemented"; return }
