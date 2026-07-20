package generator

import (
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
)

type HttpMethod struct {
	Name               string
	HTTPMethod         string
	Comment            string
	RequestTypeName    string
	RequestTypePackage string
	RequestTypeRawName string
	ReturnTypeName     string
	ReturnTypePackage  string
	ReturnTypeRawName  string
	Path               string
	Serializer         string
	OutputDir          string
	RefPackage         string
	RefPackageAlias    string
	ModelPackage       map[string]string
	GenHandler         bool
	Models             map[string]*model.Model
}

type Handler struct {
	FilePath    string
	PackageName string
	ProjPackage string
	Imports     map[string]*model.Model
	Methods     []*HttpMethod
}

type SingleHandler struct {
	*HttpMethod
	FilePath    string
	PackageName string
	ProjPackage string
}

type Client struct {
	Handler
	ServiceName string
}

func (pkgGen *HttpPackageGenerator) genHandler(pkg *HttpPackage, handlerDir, handlerPackage string, root *RouterNode) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgGen *HttpPackageGenerator) processHandler(handler *Handler, root *RouterNode, handlerDir, projectOutDir string, handlerByMethod bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgGen *HttpPackageGenerator) updateHandler(handler interface{}, handlerTpl, filePath string, noRepeat bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgGen *HttpPackageGenerator) updateClient(client interface{}, clientTpl, filePath string, noRepeat bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *HttpMethod) InitComment() { _ = "STUB: not implemented"; return }

func MapSerializer(serializer string) string { _ = "STUB: not implemented"; return "" }

func (h *Handler) Format() { _ = "STUB: not implemented"; return }
