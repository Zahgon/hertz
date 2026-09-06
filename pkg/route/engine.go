package route

import (
	"context"
	"html/template"
	"sync"
	"sync/atomic"

	"github.com/cloudwego/hertz/internal/nocopy"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/app/server/render"
	"github.com/cloudwego/hertz/pkg/common/config"
	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/tracer"
	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/network/standard"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/http1"
	"github.com/cloudwego/hertz/pkg/protocol/suite"
)

const unknownTransporterName = "unknown"

var (
	defaultTransporter = standard.NewTransporter

	errInitFailed       = errs.NewPrivate("engine has been init already")
	errAlreadyRunning   = errs.NewPrivate("engine is already running")
	errStatusNotRunning = errs.NewPrivate("engine is not running")

	default404Body = []byte("Not Found")
	default405Body = []byte("Method Not Allowed")
	default400Body = []byte("Bad Request")

	requiredHostBody = []byte("missing required Host header")
)

type hijackConn struct {
	network.Conn
	e *Engine
}

type CtxCallback func(ctx context.Context)

type CtxErrCallback func(ctx context.Context) error

type RouteInfo struct {
	Method      string
	Path        string
	Handler     string
	HandlerFunc app.HandlerFunc
}

type RoutesInfo []RouteInfo

type Engine struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	Name       string
	serverName atomic.Value

	options *config.Options

	RouterGroup
	trees MethodTrees

	maxParams uint16

	allNoMethod app.HandlersChain
	allNoRoute  app.HandlersChain
	noRoute     app.HandlersChain
	noMethod    app.HandlersChain

	delims     render.Delims
	funcMap    template.FuncMap
	htmlRender render.HTMLRender

	NoHijackConnPool bool
	hijackConnPool   sync.Pool

	KeepHijackedConns bool

	transport network.Transporter

	tracerCtl   tracer.Controller
	enableTrace bool

	protocolSuite         *suite.Config
	protocolServers       map[string]protocol.Server
	protocolStreamServers map[string]protocol.StreamServer

	ctxPool sync.Pool

	PanicHandler app.HandlerFunc

	ContinueHandler func(header *protocol.RequestHeader) bool

	status uint32

	OnRun []CtxErrCallback

	OnShutdown []CtxCallback

	clientIPFunc  app.ClientIP
	formValueFunc app.FormValueFunc

	binder binding.Binder
}

func (engine *Engine) IsTraceEnable() bool { _ = "STUB: not implemented"; return false }

func (engine *Engine) GetCtxPool() *sync.Pool { _ = "STUB: not implemented"; return nil }

func (engine *Engine) GetOptions() *config.Options { _ = "STUB: not implemented"; return nil }

func SetTransporter(transporter func(options *config.Options) network.Transporter) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) GetTransporterName() (tName string) { _ = "STUB: not implemented"; return "" }

func getTransporterName(transporter network.Transporter) (tName string) {
	_ = "STUB: not implemented"
	return ""
}

func GetTransporterName() (tName string) { _ = "STUB: not implemented"; return "" }

func (engine *Engine) IsStreamRequestBody() bool { _ = "STUB: not implemented"; return false }

func (engine *Engine) IsRunning() bool { _ = "STUB: not implemented"; return false }

func (engine *Engine) HijackConnHandle(c network.Conn, h app.HijackHandler) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) GetTracer() tracer.Controller {
	_ = "STUB: not implemented"
	return *new(tracer.Controller)
}

const (
	_ uint32 = iota
	statusInitialized
	statusRunning
	statusShutdown
	statusClosed
)

func (engine *Engine) NewContext() *app.RequestContext { _ = "STUB: not implemented"; return nil }

func (engine *Engine) Shutdown(ctx context.Context) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) executeOnShutdownHooks(ctx context.Context) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) Run() (err error) { _ = "STUB: not implemented"; return nil }

func (engine *Engine) Init() error { _ = "STUB: not implemented"; return nil }

func (engine *Engine) alpnEnable() bool { _ = "STUB: not implemented"; return false }

func (engine *Engine) listenAndServe() error { _ = "STUB: not implemented"; return nil }

func (c *hijackConn) Close() error { _ = "STUB: not implemented"; return nil }

func (engine *Engine) getNextProto(conn network.Conn) (proto string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (engine *Engine) onData(c context.Context, conn interface{}) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func logError(conn network.Conn, err error) { _ = "STUB: not implemented"; return }

func (engine *Engine) Close() error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (engine *Engine) GetServerName() []byte { _ = "STUB: not implemented"; return nil }

func (engine *Engine) Serve(c context.Context, conn network.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) ServeStream(ctx context.Context, conn network.StreamConn) error {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) initBinderAndValidator(opt *config.Options) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) initValidatorFunc(opt *config.Options) binding.ValidatorFunc {
	_ = "STUB: not implemented"
	return *new(binding.ValidatorFunc)
}

//nolint:staticcheck // Deprecated
//nolint:staticcheck // Deprecated

//nolint:staticcheck // Deprecated

//nolint:staticcheck // Deprecated

func (engine *Engine) initCustomBinder(customBinder interface{}) { _ = "STUB: not implemented"; return }

func (engine *Engine) initDefaultBinder(bindConfig interface{}, vf binding.ValidatorFunc) {
	_ = "STUB: not implemented"
	return
}

func NewEngine(opt *config.Options) *Engine { _ = "STUB: not implemented"; return nil }

func initTrace(engine *Engine) stats.Level { _ = "STUB: not implemented"; return *new(stats.Level) }

func debugPrintRoute(httpMethod, absolutePath string, handlers app.HandlersChain) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) addRoute(method, path string, handlers app.HandlersChain) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) PrintRoute(method string) { _ = "STUB: not implemented"; return }

func printNode(node *node, level int) { _ = "STUB: not implemented"; return }

func (engine *Engine) recv(ctx *app.RequestContext) { _ = "STUB: not implemented"; return }

func (engine *Engine) ServeHTTP(c context.Context, ctx *app.RequestContext) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) allocateContext() *app.RequestContext { _ = "STUB: not implemented"; return nil }

func serveError(c context.Context, ctx *app.RequestContext, code int, defaultMessage []byte) {
	_ = "STUB: not implemented"
	return
}

func trailingSlashURL(ts string) string { _ = "STUB: not implemented"; return "" }

func redirectTrailingSlash(c *app.RequestContext) { _ = "STUB: not implemented"; return }

func redirectRequest(c *app.RequestContext) { _ = "STUB: not implemented"; return }

func redirectFixedPath(c *app.RequestContext, root *node, trailingSlash bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (engine *Engine) NoRoute(handlers ...app.HandlerFunc) { _ = "STUB: not implemented"; return }

func (engine *Engine) NoMethod(handlers ...app.HandlerFunc) { _ = "STUB: not implemented"; return }

func (engine *Engine) rebuild404Handlers() { _ = "STUB: not implemented"; return }

func (engine *Engine) rebuild405Handlers() { _ = "STUB: not implemented"; return }

func (engine *Engine) Use(middleware ...app.HandlerFunc) IRoutes {
	_ = "STUB: not implemented"
	return *new(IRoutes)
}

func (engine *Engine) LoadHTMLGlob(pattern string) { _ = "STUB: not implemented"; return }

func (engine *Engine) LoadHTMLFiles(files ...string) { _ = "STUB: not implemented"; return }

func (engine *Engine) SetHTMLTemplate(tmpl *template.Template) { _ = "STUB: not implemented"; return }

func (engine *Engine) SetAutoReloadHTMLTemplate(tmpl *template.Template, files []string) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) SetFuncMap(funcMap template.FuncMap) { _ = "STUB: not implemented"; return }

func (engine *Engine) SetClientIPFunc(f app.ClientIP) { _ = "STUB: not implemented"; return }

func (engine *Engine) SetFormValueFunc(f app.FormValueFunc) { _ = "STUB: not implemented"; return }

func (engine *Engine) Delims(left, right string) *Engine { _ = "STUB: not implemented"; return nil }

func (engine *Engine) acquireHijackConn(c network.Conn) *hijackConn {
	_ = "STUB: not implemented"
	return nil
}

func (engine *Engine) releaseHijackConn(hjc *hijackConn) { _ = "STUB: not implemented"; return }

func (engine *Engine) hijackConnHandler(c network.Conn, h app.HijackHandler) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) Routes() (routes RoutesInfo) {
	_ = "STUB: not implemented"
	return *new(RoutesInfo)
}

func (engine *Engine) AddProtocol(protocol string, factory interface{}) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) SetAltHeader(targetProtocol, altHeaderValue string) {
	_ = "STUB: not implemented"
	return
}

func (engine *Engine) HasServer(name string) bool { _ = "STUB: not implemented"; return false }

func iterate(method string, routes RoutesInfo, root *node) RoutesInfo {
	_ = "STUB: not implemented"
	return *new(RoutesInfo)
}

func newHttp1OptionFromEngine(engine *Engine) *http1.Option { _ = "STUB: not implemented"; return nil }

func versionToALNP(v uint32) string { _ = "STUB: not implemented"; return "" }

func (engine *Engine) MarkAsRunning() (err error) { _ = "STUB: not implemented"; return nil }
