package http1

import (
	"context"
	"crypto/tls"
	"io"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server/render"
	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/common/tracer/traceinfo"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/suite"
)

func init() {
	if b, err := utils.GetBoolFromEnv("HERTZ_DISABLE_REQUEST_CONTEXT_POOL"); err == nil {
		disabaleRequestContextPool = b
	}
}

const NextProtoTLS = suite.HTTP1

var (
	errHijacked        = errs.New(errs.ErrHijacked, errs.ErrorTypePublic, nil)
	errIdleTimeout     = errs.New(errs.ErrIdleTimeout, errs.ErrorTypePrivate, nil)
	errShortConnection = errs.New(errs.ErrShortConnection, errs.ErrorTypePublic, "server is going to close the connection")
	errUnexpectedEOF   = errs.NewPublic(io.ErrUnexpectedEOF.Error() + " when reading request")

	disabaleRequestContextPool = false
)

type Option struct {
	StreamRequestBody             bool
	GetOnly                       bool
	NoDefaultDate                 bool
	NoDefaultContentType          bool
	DisablePreParseMultipartForm  bool
	DisableKeepalive              bool
	NoDefaultServerHeader         bool
	DisableHeaderNamesNormalizing bool
	MaxRequestBodySize            int
	MaxHeaderBytes                int
	IdleTimeout                   time.Duration
	ReadTimeout                   time.Duration
	ServerName                    []byte
	TLS                           *tls.Config
	HTMLRender                    render.HTMLRender
	EnableTrace                   bool
	ContinueHandler               func(header *protocol.RequestHeader) bool
	HijackConnHandle              func(c network.Conn, h app.HijackHandler)
}

type Server struct {
	Option
	Core suite.Core

	eventStackPool *sync.Pool
}

func (s Server) getRequestContext() *app.RequestContext { _ = "STUB: not implemented"; return nil }

func (s Server) putRequestContext(ctx *app.RequestContext) { _ = "STUB: not implemented"; return }

func (s Server) Serve(c context.Context, conn network.Conn) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func NewServer() *Server { _ = "STUB: not implemented"; return nil }

func writeErrorResponse(zw network.Writer, ctx *app.RequestContext, serverName []byte, err error) network.Writer {
	_ = "STUB: not implemented"
	return *new(network.Writer)
}

//nolint:errcheck
//nolint:errcheck

func writeResponse(ctx *app.RequestContext, w network.Writer) error {
	_ = "STUB: not implemented"
	return nil
}

func defaultErrorHandler(ctx *app.RequestContext, err error) { _ = "STUB: not implemented"; return }

type eventStack []func(ti traceinfo.TraceInfo, err error)

func (e *eventStack) isEmpty() bool { _ = "STUB: not implemented"; return false }

func (e *eventStack) push(f func(ti traceinfo.TraceInfo, err error)) {
	_ = "STUB: not implemented"
	return
}

func (e *eventStack) pop() func(ti traceinfo.TraceInfo, err error) {
	_ = "STUB: not implemented"
	return nil
}

func shouldRecordInTraceError(err error) bool { _ = "STUB: not implemented"; return false }
