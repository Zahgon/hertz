package adaptor

import (
	"bufio"
	"errors"
	"net"
	"net/http"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/network"
)

func HertzHandler(h http.Handler) app.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(app.HandlerFunc)
}

type httpResponseWriter struct {
	rc     *app.RequestContext
	header http.Header

	err error

	wroteHeader bool
	skipBody    bool

	hijacked chan struct{}
}

var errConnHijacked = errors.New("hertz net/http adaptor: conn hijacked")

func (p *httpResponseWriter) Header() http.Header {
	_ = "STUB: not implemented"
	return *new(http.Header)
}

func (p *httpResponseWriter) Write(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (p *httpResponseWriter) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

var _ http.Flusher = (*httpResponseWriter)(nil)

func (p *httpResponseWriter) Flush() { _ = "STUB: not implemented"; return }

var _ http.Hijacker = (*httpResponseWriter)(nil)

func (p *httpResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}

type noopHijackWriter struct{}

var _ network.ExtWriter = noopHijackWriter{}

func (noopHijackWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (noopHijackWriter) Flush() error    { _ = "STUB: not implemented"; return nil }
func (noopHijackWriter) Finalize() error { _ = "STUB: not implemented"; return nil }

type noopWriter struct{}

var _ network.ExtWriter = noopWriter{}

func (noopWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
func (noopWriter) Flush() error                { _ = "STUB: not implemented"; return nil }
func (noopWriter) Finalize() error             { _ = "STUB: not implemented"; return nil }

type hijackedConn struct {
	network.Conn

	closeOnce sync.Once
	closeCh   chan struct{}
}

func newHijackedConn(conn network.Conn) *hijackedConn { _ = "STUB: not implemented"; return nil }

func hijackedConnFinalizer(c *hijackedConn) { _ = "STUB: not implemented"; return }

func (c *hijackedConn) Close() error { _ = "STUB: not implemented"; return nil }
