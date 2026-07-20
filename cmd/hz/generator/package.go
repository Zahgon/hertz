package generator

import (
	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/cloudwego/hertz/cmd/hz/meta"
)

type HttpPackage struct {
	IdlName    string
	Package    string
	Services   []*Service
	Models     []*model.Model
	RouterInfo *Router
}

type Service struct {
	Name          string
	Methods       []*HttpMethod
	ClientMethods []*ClientMethod
	Models        []*model.Model
	BaseDomain    string
	ServiceGroup  string
	ServiceGenDir string
}

type HttpPackageGenerator struct {
	ConfigPath     string
	Backend        meta.Backend
	Options        []Option
	CmdType        string
	ProjPackage    string
	HandlerDir     string
	RouterDir      string
	ModelDir       string
	UseDir         string
	ClientDir      string
	IdlClientDir   string
	ForceClientDir string
	BaseDomain     string
	QueryEnumAsInt bool
	ServiceGenDir  string

	NeedModel            bool
	HandlerByMethod      bool
	SnakeStyleMiddleware bool
	SortRouter           bool
	ForceUpdateClient    bool

	loadedBackend   Backend
	curModel        *model.Model
	processedModels map[*model.Model]bool

	TemplateGenerator
}

func (pkgGen *HttpPackageGenerator) Init() error { _ = "STUB: not implemented"; return nil }

func (pkgGen *HttpPackageGenerator) checkInited() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (pkgGen *HttpPackageGenerator) Generate(pkg *HttpPackage) error {
	_ = "STUB: not implemented"
	return nil
}
