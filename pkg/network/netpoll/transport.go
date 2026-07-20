//go:build !windows

package netpoll

import (
	"context"
	"io"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/netpoll"
)

func init() {

	netpoll.SetLoggerOutput(io.Discard)
}

type ctxCancelKeyStruct struct{}

var ctxCancelKey = ctxCancelKeyStruct{}

func cancelContext(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

type transporter struct {
	senseClientDisconnection bool
	network                  string
	addr                     string
	keepAliveTimeout         time.Duration
	readTimeout              time.Duration
	writeTimeout             time.Duration
	listenConfig             *net.ListenConfig
	OnAccept                 func(conn net.Conn) context.Context
	OnConnect                func(ctx context.Context, conn network.Conn) context.Context

	mu sync.RWMutex
	ln net.Listener
	el netpoll.EventLoop
}

func NewTransporter(options *config.Options) network.Transporter {
	_ = "STUB: not implemented"
	return *new(network.Transporter)
}

func (t *transporter) Listener() net.Listener { _ = "STUB: not implemented"; return *new(net.Listener) }

func (t *transporter) ListenAndServe(onReq network.OnData) (err error) {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func (t *transporter) Close() error { _ = "STUB: not implemented"; return nil }

func (t *transporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck
