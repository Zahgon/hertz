package route

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/route/param"
)

type router struct {
	method string
	root   *node
}

type MethodTrees []*router

func (trees MethodTrees) get(method string) *router { _ = "STUB: not implemented"; return nil }

func countParams(path string) uint16 { _ = "STUB: not implemented"; return 0 }

type (
	node struct {
		kind     kind
		label    byte
		prefix   string
		parent   *node
		children children

		ppath string

		pnames     []string
		handlers   app.HandlersChain
		paramChild *node
		anyChild   *node

		isLeaf bool
	}
	kind     uint8
	children []*node
)

const (
	skind kind = iota

	pkind

	akind
	paramLabel = byte(':')
	anyLabel   = byte('*')
	slash      = "/"
	nilString  = ""
)

func checkPathValid(path string) { _ = "STUB: not implemented"; return }

func (r *router) addRoute(path string, h app.HandlersChain) { _ = "STUB: not implemented"; return }

func (r *router) insert(path string, h app.HandlersChain, t kind, ppath string, pnames []string) {
	_ = "STUB: not implemented"
	return
}

func (r *router) find(path string, paramsPointer *param.Params, unescape bool) (res nodeValue) {
	_ = "STUB: not implemented"
	return *new(nodeValue)
}

func (n *node) findChild(l byte) *node { _ = "STUB: not implemented"; return nil }

func (n *node) findChildWithLabel(l byte) *node { _ = "STUB: not implemented"; return nil }

func newNode(t kind, pre string, p *node, child children, mh app.HandlersChain, ppath string, pnames []string, paramChildren, anyChildren *node) *node {
	_ = "STUB: not implemented"
	return nil
}

type nodeValue struct {
	handlers app.HandlersChain
	tsr      bool
	fullPath string
}

func (n *node) findCaseInsensitivePath(path string, fixTrailingSlash bool) (ciPath []byte, found bool) {
	_ = "STUB: not implemented"
	return nil, false
}
