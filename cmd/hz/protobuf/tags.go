package protobuf

import (
	"github.com/cloudwego/hertz/cmd/hz/config"
	"github.com/cloudwego/hertz/cmd/hz/generator"
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/cloudwego/hertz/cmd/hz/protobuf/api"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/runtime/protoimpl"
	"google.golang.org/protobuf/types/descriptorpb"
)

var (
	jsonSnakeName             = false
	unsetOmitempty            = false
	protobufCamelJSONTagStyle = false
)

func CheckTagOption(args *config.Argument) (ret []generator.Option) {
	_ = "STUB: not implemented"
	return nil
}

func checkSnakeName(name string) string { _ = "STUB: not implemented"; return "" }

var (
	HttpMethodOptions = map[*protoimpl.ExtensionInfo]string{
		api.E_Get:     "GET",
		api.E_Post:    "POST",
		api.E_Put:     "PUT",
		api.E_Patch:   "PATCH",
		api.E_Delete:  "DELETE",
		api.E_Options: "OPTIONS",
		api.E_Head:    "HEAD",
		api.E_Any:     "Any",
	}

	BindingTags = map[*protoimpl.ExtensionInfo]string{
		api.E_Path:   "path",
		api.E_Query:  "query",
		api.E_Header: "header",
		api.E_Cookie: "cookie",
		api.E_Body:   "json",

		api.E_Form:           "form",
		api.E_FormCompatible: "form",
		api.E_RawBody:        "raw_body",
	}

	ValidatorTags = map[*protoimpl.ExtensionInfo]string{api.E_Vd: "vd"}

	SerializerOptions = map[*protoimpl.ExtensionInfo]string{api.E_Serializer: "serializer"}
)

type httpOption struct {
	method string
	path   string
}

type httpOptions []httpOption

func (s httpOptions) Len() int { _ = "STUB: not implemented"; return 0 }

func (s httpOptions) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s httpOptions) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func getAllOptions(extensions map[*protoimpl.ExtensionInfo]string, opts ...protoreflect.ProtoMessage) map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func checkFirstOptions(extensions map[*protoimpl.ExtensionInfo]string, opts ...protoreflect.ProtoMessage) (string, interface{}) {
	_ = "STUB: not implemented"
	return "", nil
}

func checkFirstOption(ext *protoimpl.ExtensionInfo, opts ...protoreflect.ProtoMessage) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func checkOption(ext *protoimpl.ExtensionInfo, opts ...protoreflect.ProtoMessage) (ret []interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func tag(k string, v interface{}) model.Tag { _ = "STUB: not implemented"; return *new(model.Tag) }

func defaultBindingTags(f *descriptorpb.FieldDescriptorProto) []model.Tag {
	_ = "STUB: not implemented"
	return nil
}

func jsonTag(f *descriptorpb.FieldDescriptorProto) (ret model.Tag) {
	_ = "STUB: not implemented"
	return *new(model.Tag)
}

func injectTagsToModel(f *descriptorpb.FieldDescriptorProto, gf *model.Field, needDefault bool) error {
	_ = "STUB: not implemented"
	return nil
}

func getJsonValue(f *descriptorpb.FieldDescriptorProto, val string) string {
	_ = "STUB: not implemented"
	return ""
}

func checkRequire(f *descriptorpb.FieldDescriptorProto, val string) string {
	_ = "STUB: not implemented"
	return ""
}

func m2s(mt model.Tag) (ret [2]string) { _ = "STUB: not implemented"; return [2]string{} }

func reflectJsonTag(f protoreflect.FieldDescriptor) (ret model.Tag) {
	_ = "STUB: not implemented"
	return *new(model.Tag)
}

func defaultBindingStructTags(f protoreflect.FieldDescriptor) []model.Tag {
	_ = "STUB: not implemented"
	return nil
}

func injectTagsToStructTags(f protoreflect.FieldDescriptor, out *structTags, needDefault bool, rmTags RemoveTags) error {
	_ = "STUB: not implemented"
	return nil
}

func getStructJsonValue(f protoreflect.FieldDescriptor, val string) string {
	_ = "STUB: not implemented"
	return ""
}

func checkStructRequire(f protoreflect.FieldDescriptor, val string) string {
	_ = "STUB: not implemented"
	return ""
}
