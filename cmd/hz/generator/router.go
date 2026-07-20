package generator

import (
	"regexp"
)

type Router struct {
	FilePath        string
	PackageName     string
	HandlerPackages map[string]string
	Router          *RouterNode
}

type RouterNode struct {
	GroupName         string
	MiddleWare        string
	HandlerMiddleware string
	GroupMiddleware   string
	PathPrefix        string

	Path     string
	Parent   *RouterNode
	Children childrenRouterInfo

	Handler             string
	HandlerPackage      string
	HandlerPackageAlias string
	HttpMethod          string
}

type RegisterInfo struct {
	PackageName string
	DepPkgAlias string
	DepPkg      string
}

func NewRouterTree() *RouterNode { _ = "STUB: not implemented"; return nil }

func (routerNode *RouterNode) Sort() { _ = "STUB: not implemented"; return }

func (routerNode *RouterNode) Update(method *HttpMethod, handlerType, handlerPkg string, sortRouter bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (routerNode *RouterNode) RawHandlerName() string { _ = "STUB: not implemented"; return "" }

func (routerNode *RouterNode) DyeGroupName(snakeStyleMiddleware bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (routerNode *RouterNode) DFS(i int, hook func(layer int, node *RouterNode) error) error {
	_ = "STUB: not implemented"
	return nil
}

var handlerPkgMap map[string]string

func (routerNode *RouterNode) Insert(name string, method *HttpMethod, handlerType string, paths []string, handlerPkg string, sortRouter bool) {
	_ = "STUB: not implemented"
	return
}

func getHttpMethod(method string) string { _ = "STUB: not implemented"; return "" }

func (routerNode *RouterNode) FindNearest(paths []string, method string, sortRouter bool) (*RouterNode, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

type childrenRouterInfo []*RouterNode

func (c childrenRouterInfo) Len() int { _ = "STUB: not implemented"; return 0 }

func (c childrenRouterInfo) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func removeNonLetterPrefix(str string) string { _ = "STUB: not implemented"; return "" }

func (c childrenRouterInfo) Swap(i, j int) { _ = "STUB: not implemented"; return }

var (
	regRegisterV3 = regexp.MustCompile(insertPointPatternNew)
	regImport     = regexp.MustCompile(`import \(\n`)
)

func (pkgGen *HttpPackageGenerator) updateRegister(pkg, rDir, pkgName string) error {
	_ = "STUB: not implemented"
	return nil
}

func checkDupRegister(file []byte, insertReg string) bool { _ = "STUB: not implemented"; return false }

func appendMw(mws []string, mw string) ([]string, string) {
	_ = "STUB: not implemented"
	return nil, ""
}

func stringsIncludes(strs []string, str string) bool { _ = "STUB: not implemented"; return false }

func (pkgGen *HttpPackageGenerator) genRouter(pkg *HttpPackage, root *RouterNode, handlerPackage, routerDir, routerPackage string) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgGen *HttpPackageGenerator) updateMiddlewareReg(router interface{}, middlewareTpl, filePath string) error {
	_ = "STUB: not implemented"
	return nil
}

func convertToMiddlewareName(path string) string { _ = "STUB: not implemented"; return "" }
