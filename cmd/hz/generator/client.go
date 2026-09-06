package generator

import (
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
)

type ClientMethod struct {
	*HttpMethod
	BodyParamsCode   string
	QueryParamsCode  string
	PathParamsCode   string
	HeaderParamsCode string
	FormValueCode    string
	FormFileCode     string
}

type ClientConfig struct {
	QueryEnumAsInt bool
}

type ClientFile struct {
	Config        ClientConfig
	FilePath      string
	PackageName   string
	ServiceName   string
	BaseDomain    string
	Imports       map[string]*model.Model
	ClientMethods []*ClientMethod
}

func (pkgGen *HttpPackageGenerator) genClient(pkg *HttpPackage, clientDir string) error {
	_ = "STUB: not implemented"
	return nil
}
