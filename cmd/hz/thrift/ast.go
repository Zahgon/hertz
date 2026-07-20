package thrift

import (
	"github.com/cloudwego/hertz/cmd/hz/config"
	"github.com/cloudwego/hertz/cmd/hz/generator"
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/cloudwego/thriftgo/parser"
)

func getGoPackage(ast *parser.Thrift, pkgMap map[string]string) string {
	_ = "STUB: not implemented"
	return ""
}

func astToService(ast *parser.Thrift, resolver *Resolver, args *config.Argument) ([]*generator.Service, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func newHTTPMethod(s *parser.Service, m *parser.Function, method *generator.HttpMethod, i int, anno httpAnnotation) (*generator.HttpMethod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func parseAnnotationToClient(clientMethod *generator.ClientMethod, p *parser.Type, symbol ResolvedSymbol, enableOptional bool) error {
	_ = "STUB: not implemented"
	return nil
}

type extendServiceList []string

func (svr extendServiceList) exist(serviceName string) bool {
	_ = "STUB: not implemented"
	return false
}

func getExtendServices(ast *parser.Thrift) (res extendServiceList) {
	_ = "STUB: not implemented"
	return *new(extendServiceList)
}

func getAllExtendFunction(svc *parser.Service, ast *parser.Thrift, resolver *Resolver, args *config.Argument) (res []*parser.Function, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func processExtendsType(f *parser.Function, base string) { _ = "STUB: not implemented"; return }

func getUniqueResolveDependentName(name string, resolver *Resolver) string {
	_ = "STUB: not implemented"
	return ""
}

func addResolverDependency(resolver *Resolver, ast *parser.Thrift, args *config.Argument) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

var BaseThrift = parser.Thrift{}

var baseTypes = map[string]string{
	"bool":   "bool",
	"byte":   "int8",
	"i8":     "int8",
	"i16":    "int16",
	"i32":    "int32",
	"i64":    "int64",
	"double": "float64",
	"string": "string",
	"binary": "[]byte",
}

func switchBaseType(typ *parser.Type) *model.Type { _ = "STUB: not implemented"; return nil }

func newBaseType(typ *model.Type, cg model.Category) *model.Type {
	_ = "STUB: not implemented"
	return nil
}

func newStructType(name string, cg model.Category) *model.Type {
	_ = "STUB: not implemented"
	return nil
}

func newEnumType(name string, cg model.Category) *model.Type { _ = "STUB: not implemented"; return nil }

func newFuncType(name string, cg model.Category) *model.Type { _ = "STUB: not implemented"; return nil }

func (resolver *Resolver) getFieldType(typ *parser.Type) (*model.Type, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type ResolvedSymbol struct {
	Base string
	Src  string
	*Symbol
}

func (rs ResolvedSymbol) Expression() string { _ = "STUB: not implemented"; return "" }

func astToModel(ast *parser.Thrift, rs *Resolver) (*model.Model, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func removeCommentsSlash(comments string) string { _ = "STUB: not implemented"; return "" }

func isPointer(f *parser.Field, rs *Resolver) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func getNewFieldName(fieldName string, fieldNameSet map[string]bool) string {
	_ = "STUB: not implemented"
	return ""
}

func checkDuplicatedFileName(vs []model.Field) { _ = "STUB: not implemented"; return }
