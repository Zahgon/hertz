package thrift

import (
	"github.com/cloudwego/hertz/cmd/hz/config"
	"github.com/cloudwego/hertz/cmd/hz/generator"
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/cloudwego/thriftgo/parser"
)

const (
	AnnotationQuery    = "api.query"
	AnnotationForm     = "api.form"
	AnnotationPath     = "api.path"
	AnnotationHeader   = "api.header"
	AnnotationCookie   = "api.cookie"
	AnnotationBody     = "api.body"
	AnnotationRawBody  = "api.raw_body"
	AnnotationJsConv   = "api.js_conv"
	AnnotationNone     = "api.none"
	AnnotationFileName = "api.file_name"

	AnnotationValidator = "api.vd"

	AnnotationGoTag = "go.tag"
)

const (
	ApiGet        = "api.get"
	ApiPost       = "api.post"
	ApiPut        = "api.put"
	ApiPatch      = "api.patch"
	ApiDelete     = "api.delete"
	ApiOptions    = "api.options"
	ApiHEAD       = "api.head"
	ApiAny        = "api.any"
	ApiPath       = "api.path"
	ApiSerializer = "api.serializer"
	ApiGenPath    = "api.handler_path"
)

const (
	ApiBaseDomain    = "api.base_domain"
	ApiServiceGroup  = "api.service_group"
	ApiServiceGenDir = "api.service_gen_dir"
	ApiServicePath   = "api.service_path"
)

var (
	HttpMethodAnnotations = map[string]string{
		ApiGet:     "GET",
		ApiPost:    "POST",
		ApiPut:     "PUT",
		ApiPatch:   "PATCH",
		ApiDelete:  "DELETE",
		ApiOptions: "OPTIONS",
		ApiHEAD:    "HEAD",
		ApiAny:     "ANY",
	}

	HttpMethodOptionAnnotations = map[string]string{
		ApiGenPath: "handler_path",
	}

	BindingTags = map[string]string{
		AnnotationPath:    "path",
		AnnotationQuery:   "query",
		AnnotationHeader:  "header",
		AnnotationCookie:  "cookie",
		AnnotationBody:    "json",
		AnnotationForm:    "form",
		AnnotationRawBody: "raw_body",
	}

	SerializerTags = map[string]string{
		ApiSerializer: "serializer",
	}

	ValidatorTags = map[string]string{AnnotationValidator: "vd"}
)

var (
	jsonSnakeName  = false
	unsetOmitempty = false
)

func CheckTagOption(args *config.Argument) []generator.Option {
	_ = "STUB: not implemented"
	return nil
}

func checkSnakeName(name string) string { _ = "STUB: not implemented"; return "" }

func getAnnotation(input parser.Annotations, target string) []string {
	_ = "STUB: not implemented"
	return nil
}

type httpAnnotation struct {
	method string
	path   []string
}

type httpAnnotations []httpAnnotation

func (s httpAnnotations) Len() int { _ = "STUB: not implemented"; return 0 }

func (s httpAnnotations) Swap(i, j int) { _ = "STUB: not implemented"; return }

func (s httpAnnotations) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func getAnnotations(input parser.Annotations, targets map[string]string) map[string][]string {
	_ = "STUB: not implemented"
	return nil
}

func defaultBindingTags(f *parser.Field) []model.Tag { _ = "STUB: not implemented"; return nil }

func jsonTag(f *parser.Field) (ret model.Tag) { _ = "STUB: not implemented"; return *new(model.Tag) }

func tag(k, v string) model.Tag { _ = "STUB: not implemented"; return *new(model.Tag) }

func annotationToTags(as parser.Annotations, targets map[string]string) (tags []model.Tag) {
	_ = "STUB: not implemented"
	return nil
}

func injectTags(f *parser.Field, gf *model.Field, needDefault, needGoTag bool) error {
	_ = "STUB: not implemented"
	return nil
}

func getJsonValue(f *parser.Field, val string) string { _ = "STUB: not implemented"; return "" }

func checkRequire(f *parser.Field, val string) string { _ = "STUB: not implemented"; return "" }

func checkGoTag(as parser.Annotations, tags *model.Tags) error {
	_ = "STUB: not implemented"
	return nil
}
