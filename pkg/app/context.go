package app

import (
	"context"
	"io"
	"mime/multipart"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/app/server/render"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/tracer/traceinfo"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/route/param"
)

var zeroTCPAddr = &net.TCPAddr{
	IP: net.IPv4zero,
}

type Handler interface {
	ServeHTTP(c context.Context, ctx *RequestContext)
}

type ClientIP func(ctx *RequestContext) string

type ClientIPOptions struct {
	RemoteIPHeaders []string
	TrustedCIDRs    []*net.IPNet
}

var defaultTrustedCIDRs = []*net.IPNet{
	{
		IP:   net.IP{0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0},
	},
	{
		IP:   net.IP{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
		Mask: net.IPMask{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
	},
}

var defaultClientIPOptions = ClientIPOptions{
	RemoteIPHeaders: []string{"X-Forwarded-For", "X-Real-IP"},
	TrustedCIDRs:    defaultTrustedCIDRs,
}

var loopbackIP = net.ParseIP("127.0.0.1")

func ClientIPWithOption(opts ClientIPOptions) ClientIP {
	_ = "STUB: not implemented"
	return *new(ClientIP)
}

func isTrustedProxy(trustedCIDRs []*net.IPNet, remoteIP net.IP) bool {
	_ = "STUB: not implemented"
	return false
}

func validateHeader(trustedCIDRs []*net.IPNet, header string) (clientIP string, valid bool) {
	_ = "STUB: not implemented"
	return "", false
}

var defaultClientIP = ClientIPWithOption(defaultClientIPOptions)

func SetClientIPFunc(fn ClientIP) { _ = "STUB: not implemented"; return }

type FormValueFunc func(*RequestContext, string) []byte

var defaultFormValue = func(ctx *RequestContext, key string) []byte {
	v := ctx.QueryArgs().Peek(key)
	if len(v) > 0 {
		return v
	}
	v = ctx.PostArgs().Peek(key)
	if len(v) > 0 {
		return v
	}
	mf, err := ctx.MultipartForm()
	if err == nil && mf.Value != nil {
		vv := mf.Value[key]
		if len(vv) > 0 {
			return []byte(vv[0])
		}
	}
	return nil
}

type RequestContext struct {
	conn     network.Conn
	Request  protocol.Request
	Response protocol.Response

	Errors errors.ErrorChain

	Params     param.Params
	handlers   HandlersChain
	fullPath   string
	index      int8
	HTMLRender render.HTMLRender

	mu sync.RWMutex

	Keys map[string]interface{}

	hijackHandler HijackHandler

	finishedMu sync.Mutex

	finished chan struct{}

	traceInfo traceinfo.TraceInfo

	enableTrace bool

	clientIPFunc ClientIP

	formValueFunc FormValueFunc

	binder binding.Binder
	exiled bool
}

func (ctx *RequestContext) Exile() { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) IsExiled() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) Flush() error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) SetClientIPFunc(f ClientIP) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) SetFormValueFunc(f FormValueFunc) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) SetBinder(binder binding.Binder) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) GetTraceInfo() traceinfo.TraceInfo {
	_ = "STUB: not implemented"
	return *new(traceinfo.TraceInfo)
}

func (ctx *RequestContext) SetTraceInfo(t traceinfo.TraceInfo) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) IsEnableTrace() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) SetEnableTrace(enable bool) { _ = "STUB: not implemented"; return }

func NewContext(maxParams uint16) *RequestContext { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) ForEachKey(fn func(k string, v interface{})) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) SetConn(c network.Conn) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) GetConn() network.Conn {
	_ = "STUB: not implemented"
	return *new(network.Conn)
}

func (ctx *RequestContext) SetHijackHandler(h HijackHandler) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) GetHijackHandler() HijackHandler {
	_ = "STUB: not implemented"
	return *new(HijackHandler)
}

func (ctx *RequestContext) GetReader() network.Reader {
	_ = "STUB: not implemented"
	return *new(network.Reader)
}

func (ctx *RequestContext) GetWriter() network.Writer {
	_ = "STUB: not implemented"
	return *new(network.Writer)
}

func (ctx *RequestContext) GetIndex() int8 { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) SetIndex(index int8) { _ = "STUB: not implemented"; return }

type HandlerFunc func(c context.Context, ctx *RequestContext)

type HandlersChain []HandlerFunc

type HandlerNameOperator interface {
	SetHandlerName(handler HandlerFunc, name string)
	GetHandlerName(handler HandlerFunc) string
}

func SetHandlerNameOperator(o HandlerNameOperator) { _ = "STUB: not implemented"; return }

type inbuiltHandlerNameOperatorStruct struct {
	handlerNames map[uintptr]string
}

func (o *inbuiltHandlerNameOperatorStruct) SetHandlerName(handler HandlerFunc, name string) {
	_ = "STUB: not implemented"
	return
}

func (o *inbuiltHandlerNameOperatorStruct) GetHandlerName(handler HandlerFunc) string {
	_ = "STUB: not implemented"
	return ""
}

type concurrentHandlerNameOperatorStruct struct {
	handlerNames map[uintptr]string
	lock         sync.RWMutex
}

func (o *concurrentHandlerNameOperatorStruct) SetHandlerName(handler HandlerFunc, name string) {
	_ = "STUB: not implemented"
	return
}

func (o *concurrentHandlerNameOperatorStruct) GetHandlerName(handler HandlerFunc) string {
	_ = "STUB: not implemented"
	return ""
}

func SetConcurrentHandlerNameOperator() { _ = "STUB: not implemented"; return }

func init() {
	inbuiltHandlerNameOperator = &inbuiltHandlerNameOperatorStruct{handlerNames: map[uintptr]string{}}
}

var inbuiltHandlerNameOperator HandlerNameOperator

func SetHandlerName(handler HandlerFunc, name string) { _ = "STUB: not implemented"; return }

func GetHandlerName(handler HandlerFunc) string { _ = "STUB: not implemented"; return "" }

func getFuncAddr(v interface{}) uintptr { _ = "STUB: not implemented"; return 0 }

type HijackHandler func(c network.Conn)

func (ctx *RequestContext) Hijack(handler HijackHandler) { _ = "STUB: not implemented"; return }

func (c HandlersChain) Last() HandlerFunc { _ = "STUB: not implemented"; return *new(HandlerFunc) }

func (ctx *RequestContext) Finished() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) GetRequest() (dst *protocol.Request) {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) GetResponse() (dst *protocol.Response) {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) Value(key interface{}) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) Hijacked() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) SetBodyStream(bodyStream io.Reader, bodySize int) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) Host() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (ctx *RequestContext) WriteString(s string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (ctx *RequestContext) SetContentType(contentType string) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Path() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) NotModified() { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) IfModifiedSince(lastModified time.Time) bool {
	_ = "STUB: not implemented"
	return false
}

func (ctx *RequestContext) URI() *protocol.URI { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) String(code int, format string, values ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) FullPath() string { _ = "STUB: not implemented"; return "" }

func (ctx *RequestContext) SetFullPath(p string) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (ctx *RequestContext) File(filepath string) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) FileFromFS(filepath string, fs *FS) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) FileAttachment(filepath, filename string) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) SetBodyString(body string) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) SetContentTypeBytes(contentType []byte) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) FormFile(name string) (*multipart.FileHeader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *RequestContext) FormValue(key string) []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) multipartFormValue(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ctx *RequestContext) multipartFormValueArray(key string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (ctx *RequestContext) RequestBodyStream() io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (ctx *RequestContext) MultipartForm() (*multipart.Form, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ctx *RequestContext) SaveUploadedFile(file *multipart.FileHeader, dst string) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) SetConnectionClose() { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) IsGet() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) IsHead() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) IsPost() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) Method() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) NotFound() { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) redirect(uri []byte, statusCode int) { _ = "STUB: not implemented"; return }

func getRedirectStatusCode(statusCode int) int { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) Copy() *RequestContext { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) Next(c context.Context) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Handler() HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func (ctx *RequestContext) Handlers() HandlersChain {
	_ = "STUB: not implemented"
	return *new(HandlersChain)
}

func (ctx *RequestContext) SetHandlers(hc HandlersChain) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) HandlerName() string { _ = "STUB: not implemented"; return "" }

func (ctx *RequestContext) ResetWithoutConn() { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Reset() { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Redirect(statusCode int, uri []byte) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Header(key, value string) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Set(key string, value interface{}) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) Get(key string) (value interface{}, exists bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (ctx *RequestContext) MustGet(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) GetString(key string) (s string) { _ = "STUB: not implemented"; return "" }

func (ctx *RequestContext) GetBool(key string) (b bool) { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) GetInt(key string) (i int) { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) GetInt32(key string) (i32 int32) { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) GetInt64(key string) (i64 int64) { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) GetUint(key string) (ui uint) { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) GetUint32(key string) (ui32 uint32) { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) GetUint64(key string) (ui64 uint64) { _ = "STUB: not implemented"; return 0 }

func (ctx *RequestContext) GetFloat32(key string) (f32 float32) {
	_ = "STUB: not implemented"
	return 0
}

func (ctx *RequestContext) GetFloat64(key string) (f64 float64) {
	_ = "STUB: not implemented"
	return 0
}

func (ctx *RequestContext) GetTime(key string) (t time.Time) {
	_ = "STUB: not implemented"
	return *new(time.Time)
}

func (ctx *RequestContext) GetDuration(key string) (d time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (ctx *RequestContext) GetStringSlice(key string) (ss []string) {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) GetStringMap(key string) (sm map[string]interface{}) {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) GetStringMapString(key string) (sms map[string]string) {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) GetStringMapStringSlice(key string) (smss map[string][]string) {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) Param(key string) string { _ = "STUB: not implemented"; return "" }

func (ctx *RequestContext) Abort() { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) AbortWithStatus(code int) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) AbortWithMsg(msg string, statusCode int) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) AbortWithStatusJSON(code int, jsonObj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) Render(code int, r render.Render) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) ProtoBuf(code int, obj interface{}) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) JSON(code int, obj interface{}) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) PureJSON(code int, obj interface{}) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) IndentedJSON(code int, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) HTML(code int, name string, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) Data(code int, contentType string, data []byte) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) XML(code int, obj interface{}) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) AbortWithError(code int, err error) *errors.Error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) IsAborted() bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) Error(err error) *errors.Error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) ContentType() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) Cookie(key string) []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) SetCookie(name, value string, maxAge int, path, domain string, sameSite protocol.CookieSameSite, secure, httpOnly bool) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) setCookie(name, value string, maxAge int, path, domain string, sameSite protocol.CookieSameSite, secure, httpOnly, partitioned bool) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) SetPartitionedCookie(name, value string, maxAge int, path, domain string, sameSite protocol.CookieSameSite, secure, httpOnly bool) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) UserAgent() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) Status(code int) { _ = "STUB: not implemented"; return }

func (ctx *RequestContext) GetHeader(key string) []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) GetRawData() []byte { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) Body() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (ctx *RequestContext) ClientIP() string { _ = "STUB: not implemented"; return "" }

func (ctx *RequestContext) QueryArgs() *protocol.Args { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) PostArgs() *protocol.Args { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) Query(key string) string { _ = "STUB: not implemented"; return "" }

func (ctx *RequestContext) DefaultQuery(key, defaultValue string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ctx *RequestContext) GetQuery(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ctx *RequestContext) PostForm(key string) string { _ = "STUB: not implemented"; return "" }

func (ctx *RequestContext) PostFormArray(key string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) DefaultPostForm(key, defaultValue string) string {
	_ = "STUB: not implemented"
	return ""
}

func (ctx *RequestContext) GetPostForm(key string) (string, bool) {
	_ = "STUB: not implemented"
	return "", false
}

func (ctx *RequestContext) GetPostFormArray(key string) ([]string, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func bodyAllowedForStatus(status int) bool { _ = "STUB: not implemented"; return false }

func (ctx *RequestContext) getBinder() binding.Binder {
	_ = "STUB: not implemented"
	return *new(binding.Binder)
}

func (ctx *RequestContext) BindAndValidate(obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) Bind(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) Validate(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) BindQuery(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) BindHeader(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) BindPath(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) BindForm(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) BindJSON(obj interface{}) error { _ = "STUB: not implemented"; return nil }

func (ctx *RequestContext) BindProtobuf(obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) BindByContentType(obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (ctx *RequestContext) VisitAllQueryArgs(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) VisitAllPostArgs(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) VisitAllHeaders(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (ctx *RequestContext) VisitAllCookie(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}
