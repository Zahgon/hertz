package golang

import (
	"text/template"
)

var tpls *template.Template

var list = map[string]string{
	"file":      file,
	"typedef":   typedef,
	"constants": constants,
	"variables": variables,
	"function":  function,
	"enum":      enum,
	"struct":    structLike,
	"method":    method,
	"oneof":     oneof,
}

func Template() (*template.Template, error) { _ = "STUB: not implemented"; return nil, nil }

func List() map[string]string { _ = "STUB: not implemented"; return nil }

var funcMap = template.FuncMap{
	"Features":            getFeatures,
	"Identify":            identify,
	"CamelCase":           camelCase,
	"SnakeCase":           snakeCase,
	"GetTypedefReturnStr": getTypedefReturnStr,
}

func Funcs(name string, fn interface{}) error { _ = "STUB: not implemented"; return nil }

func identify(name string) string { _ = "STUB: not implemented"; return "" }

func camelCase(name string) string { _ = "STUB: not implemented"; return "" }

func snakeCase(name string) string { _ = "STUB: not implemented"; return "" }

func getTypedefReturnStr(name string) string { _ = "STUB: not implemented"; return "" }

type feature struct {
	MarshalEnumToText  bool
	TypedefAsTypeAlias bool
}

var features = feature{}

func getFeatures() feature { _ = "STUB: not implemented"; return *new(feature) }

func SetOption(opt string) error { _ = "STUB: not implemented"; return nil }

var Options = []string{
	"MarshalEnumToText",
	"TypedefAsTypeAlias",
}

func GetOptions() []string { _ = "STUB: not implemented"; return nil }
