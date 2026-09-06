package standard

import (
	"context"
	"crypto/tls"
	"errors"
	"net"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/network"
)

type transport struct {
	readBufferSize           int
	network                  string
	addr                     string
	keepAliveTimeout         time.Duration
	senseClientDisconnection bool
	readTimeout              time.Duration
	handler                  network.OnData
	tls                      *tls.Config
	listenConfig             *net.ListenConfig
	OnAccept                 func(conn net.Conn) context.Context
	OnConnect                func(ctx context.Context, conn network.Conn) context.Context

	active       int32
	shuttingDown int32

	mu sync.RWMutex
	ln net.Listener
}

func (t *transport) Listener() net.Listener { _ = "STUB: not implemented"; return *new(net.Listener) }

func (t *transport) serve() (err error) { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func (t *transport) updateActive(delta int32) int32 { _ = "STUB: not implemented"; return 0 }

func (t *transport) ListenAndServe(onData network.OnData) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (t *transport) Close() error { _ = "STUB: not implemented"; return nil }

var (
	shutdownTimeout = 30 * time.Second
	shutdownTicker  = 10 * time.Millisecond

	errShutdownTimeout = errors.New("shutdown timeout")
)

func (t *transport) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

func NewTransporter(options *config.Options) network.Transporter {
	_ = "STUB: not implemented"
	return *new(network.Transporter)
}
