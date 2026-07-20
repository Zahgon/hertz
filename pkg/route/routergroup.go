package route

import (
	"regexp"

	"github.com/cloudwego/hertz/pkg/app"
)

type IRouter interface {
	IRoutes
	Group(string, ...app.HandlerFunc) *RouterGroup
}

type IRoutes interface {
	Use(...app.HandlerFunc) IRoutes
	Handle(string, string, ...app.HandlerFunc) IRoutes
	Any(string, ...app.HandlerFunc) IRoutes
	GET(string, ...app.HandlerFunc) IRoutes
	POST(string, ...app.HandlerFunc) IRoutes
	DELETE(string, ...app.HandlerFunc) IRoutes
	PATCH(string, ...app.HandlerFunc) IRoutes
	PUT(string, ...app.HandlerFunc) IRoutes
	OPTIONS(string, ...app.HandlerFunc) IRoutes
	HEAD(string, ...app.HandlerFunc) IRoutes
	StaticFile(string, string) IRoutes
	Static(string, string) IRoutes
	StaticFS(string, *app.FS) IRoutes
}

type RouterGroup struct {
	Handlers app.HandlersChain
	basePath string
	engine   *Engine
	root     bool
}

var _ IRouter = (*RouterGroup)(nil)

func (group *RouterGroup) Use(middleware ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Group(relativePath string, handlers ...app.HandlerFunc) *RouterGroup {
	_ = "STUB: not implemented"
	return nil
}

func (group *RouterGroup) BasePath() string { _ = "STUB: not implemented"; return "" }

func (group *RouterGroup) handle(httpMethod, relativePath string, handlers app.HandlersChain) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

var upperLetterReg = regexp.MustCompile("^[A-Z]+$")

func (group *RouterGroup) Handle(httpMethod, relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) POST(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) GET(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) DELETE(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) PATCH(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) PUT(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) OPTIONS(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) HEAD(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Any(relativePath string, handlers ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) StaticFile(relativePath, filepath string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) Static(relativePath, root string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) StaticFS(relativePath string, fs *app.FS) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) combineHandlers(handlers app.HandlersChain) app.HandlersChain {
	_ = "STUB: not implemented"
	return *new(app.HandlersChain)
}

func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	_ = "STUB: not implemented"
	return ""
}

func (group *RouterGroup) returnObj() IRoutes { _ = "STUB: not implemented"; return *new(IRoutes) }

func (group *RouterGroup) GETEX(relativePath string, handler app.HandlerFunc, handlerName string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) POSTEX(relativePath string, handler app.HandlerFunc, handlerName string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) PUTEX(relativePath string, handler app.HandlerFunc, handlerName string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) DELETEEX(relativePath string, handler app.HandlerFunc, handlerName string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) HEADEX(relativePath string, handler app.HandlerFunc, handlerName string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) AnyEX(relativePath string, handler app.HandlerFunc, handlerName string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (group *RouterGroup) HandleEX(httpMethod, relativePath string, handler app.HandlerFunc, handlerName string) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func joinPaths(absolutePath, relativePath string) string { _ = "STUB: not implemented"; return "" }

func lastChar(str string) uint8 { _ = "STUB: not implemented"; return 0 }
