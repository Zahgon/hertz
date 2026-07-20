package protocol

import (
	"sync"
	"sync/atomic"

	"github.com/cloudwego/hertz/internal/nocopy"
)

var (
	ServerDate     atomic.Value
	ServerDateOnce sync.Once
)

type RequestHeader struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	disableNormalizing   bool
	connectionClose      bool
	noDefaultContentType bool

	cookiesCollected bool

	contentLength      int
	contentLengthBytes []byte

	method      []byte
	requestURI  []byte
	host        []byte
	contentType []byte

	userAgent []byte
	mulHeader [][]byte
	protocol  string

	h       []argsKV
	bufKV   argsKV
	trailer *Trailer

	cookies []argsKV

	rawHeaders []byte
}

func (h *RequestHeader) SetRawHeaders(r []byte) { _ = "STUB: not implemented"; return }

type ResponseHeader struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	disableNormalizing   bool
	connectionClose      bool
	noDefaultContentType bool
	noDefaultDate        bool

	statusCode         int
	contentLength      int
	contentLengthBytes []byte
	contentEncoding    []byte

	contentType []byte
	server      []byte
	mulHeader   [][]byte
	protocol    string

	h       []argsKV
	bufKV   argsKV
	trailer *Trailer

	cookies []argsKV

	headerLength int
}

func (h *ResponseHeader) SetHeaderLength(length int) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) GetHeaderLength() int { _ = "STUB: not implemented"; return 0 }

func (h *ResponseHeader) SetContentRange(startPos, endPos, contentLength int) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) NoDefaultContentType() bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) SetConnectionClose(close bool) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) PeekArgBytes(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetNoHTTP11(b bool) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Cookie(cookie *Cookie) bool { _ = "STUB: not implemented"; return false }

//nolint:errcheck

func (h *ResponseHeader) FullCookie() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) IsHTTP11() bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) SetContentType(contentType string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) GetHeaders() []argsKV { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) Reset() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) CopyTo(dst *ResponseHeader) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Add(key, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) VisitAll(f func(key, value []byte)) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) IsHTTP11() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) SetProtocol(p string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) GetProtocol() string { _ = "STUB: not implemented"; return "" }

func (h *RequestHeader) SetNoHTTP11(b bool) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) InitBufValue(size int) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) GetBufValue() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) HasAcceptEncodingBytes(acceptEncoding []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *RequestHeader) PeekIfModifiedSinceBytes() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) RequestURI() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekArgBytes(key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) RawHeaders() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Header() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) IsPut() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsHead() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsPost() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsDelete() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsConnect() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IgnoreBody() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) ContentLength() int { _ = "STUB: not implemented"; return 0 }

func (h *RequestHeader) SetHost(host string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func checkWriteHeaderCode(code int) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) ResetSkipNormalize() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) ContentLength() int { _ = "STUB: not implemented"; return 0 }

func (h *ResponseHeader) Set(key, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Add(key, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetContentLength(contentLength int) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) ContentLengthBytes() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) InitContentLengthWithValue(contentLength int) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) VisitAllCookie(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) DelAllCookies() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelCookie(key string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelCookieBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Header() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) PeekLocation() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) DelClientCookie(key string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DelClientCookieBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) IsDisableNormalizing() bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) ParseSetCookie(value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) PeekAll(key string) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) peekAll(key []byte) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekAll(key string) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) peekAll(key []byte) [][]byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetContentTypeBytes(contentType []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) ContentEncoding() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetContentEncoding(contentEncoding string) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) SetContentEncodingBytes(contentEncoding []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) SetContentLengthBytes(contentLength []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) SetCanonical(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) ResetConnectionClose() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Server() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) AddArgBytes(key, value []byte, noValue bool) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) SetArgBytes(key, value []byte, noValue bool) {
	_ = "STUB: not implemented"
	return
}

func (h *ResponseHeader) AppendBytes(dst []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) ConnectionClose() bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) GetCookies() []argsKV { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) ContentType() []byte { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetNoDefaultContentType(b bool) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetNoDefaultDate(b bool) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetServerBytes(server []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) MustSkipContentLength() bool { _ = "STUB: not implemented"; return false }

func (h *ResponseHeader) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (h *ResponseHeader) Del(key string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) del(key []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetBytesV(key string, value []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *RequestHeader) Len() int { _ = "STUB: not implemented"; return 0 }

func (h *RequestHeader) Reset() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetByteRange(startPos, endPos int) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) DelBytes(key []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Del(key string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetArgBytes(key, value []byte, noValue bool) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) del(key []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) CopyTo(dst *RequestHeader) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetMultipartFormBoundary(boundary string) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) ContentLengthBytes() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetContentLengthBytes(contentLength []byte) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetContentTypeBytes(contentType []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) ContentType() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) SetNoDefaultContentType(b bool) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetContentLength(contentLength int) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) InitContentLengthWithValue(contentLength int) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) MultipartFormBoundary() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) ConnectionClose() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) Method() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) IsGet() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsOptions() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) IsTrace() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) SetHostBytes(host []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetRequestURIBytes(requestURI []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetBytesKV(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) AddArgBytes(key, value []byte, noValue bool) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetUserAgentBytes(userAgent []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetCookie(key, value string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) SetCookie(cookie *Cookie) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Cookie(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Cookies() []*Cookie { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekRange() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) PeekContentEncoding() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) FullCookie() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) DelCookie(key string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) DelAllCookies() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) VisitAllCookie(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) collectCookies() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetConnectionClose(close bool) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) ResetConnectionClose() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetMethod(method string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetRequestURI(requestURI string) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) Set(key, value string) { _ = "STUB: not implemented"; return }

func initHeaderKV(kv *argsKV, key, value string, disableNormalizing bool) {
	_ = "STUB: not implemented"
	return
}

func (h *RequestHeader) SetCanonical(key, value []byte) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) ResetSkipNormalize() { _ = "STUB: not implemented"; return }

func peekRawHeader(buf, key []byte) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Host() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) UserAgent() []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) DisableNormalizing() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) IsDisableNormalizing() bool { _ = "STUB: not implemented"; return false }

func (h *RequestHeader) String() string { _ = "STUB: not implemented"; return "" }

func (h *RequestHeader) VisitAll(f func(key, value []byte)) { _ = "STUB: not implemented"; return }

func (h *RequestHeader) VisitAllCustomHeader(f func(key, value []byte)) {
	_ = "STUB: not implemented"
	return
}

func ParseContentLength(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func appendArgBytes(args []argsKV, key, value []byte, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func appendArg(args []argsKV, key, value string, noValue bool) []argsKV {
	_ = "STUB: not implemented"
	return nil
}

func (h *RequestHeader) peek(key string) []byte { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (h *ResponseHeader) Get(key string) string { _ = "STUB: not implemented"; return "" }

func (h *RequestHeader) GetAll(key string) []string { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) GetAll(key string) []string { _ = "STUB: not implemented"; return nil }

func appendHeaderLine(dst, key, value []byte) []byte { _ = "STUB: not implemented"; return nil }

func appendHeaderValue(dst, v []byte) []byte { _ = "STUB: not implemented"; return nil }

func UpdateServerDate() { _ = "STUB: not implemented"; return }

func refreshServerDate() { _ = "STUB: not implemented"; return }

func (h *RequestHeader) SetMethodBytes(method []byte) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) DisableNormalizing() { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) setSpecialHeader(key, value []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *RequestHeader) setSpecialHeader(key, value []byte) bool {
	_ = "STUB: not implemented"
	return false
}

func (h *ResponseHeader) Trailer() *Trailer { _ = "STUB: not implemented"; return nil }

func (h *RequestHeader) Trailer() *Trailer { _ = "STUB: not implemented"; return nil }

func (h *ResponseHeader) SetProtocol(p string) { _ = "STUB: not implemented"; return }

func (h *ResponseHeader) GetProtocol() string { _ = "STUB: not implemented"; return "" }
