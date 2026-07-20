package suite

import (
	"context"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/tracer"
	"github.com/cloudwego/hertz/pkg/protocol"
)

const (
	HTTP1 = "http/1.1"
	HTTP2 = "h2"

	HTTP3Draft29 = "h3-29"

	HTTP3 = "h3"
)

type Core interface {
	IsRunning() bool

	GetCtxPool() *sync.Pool

	ServeHTTP(c context.Context, ctx *app.RequestContext)

	GetTracer() tracer.Controller
}

type ServerFactory interface {
	New(core Core) (server protocol.Server, err error)
}

type StreamServerFactory interface {
	New(core Core) (server protocol.StreamServer, err error)
}

type Config struct {
	altServerConfig *altServerConfig
	configMap       map[string]ServerFactory
	streamConfigMap map[string]StreamServerFactory
}

type ServerMap map[string]protocol.Server

type StreamServerMap map[string]protocol.StreamServer

type altServerConfig struct {
	targetProtocol   string
	setAltHeaderFunc func(ctx context.Context, reqCtx *app.RequestContext)
}

type coreWrapper struct {
	Core
	beforeHandler func(c context.Context, ctx *app.RequestContext)
}

func (c *coreWrapper) ServeHTTP(ctx context.Context, reqCtx *app.RequestContext) {
	_ = "STUB: not implemented"
	return
}

func (c *Config) SetAltHeader(target, altHeader string) { _ = "STUB: not implemented"; return }

func (c *Config) Add(protocol string, factory interface{}) { _ = "STUB: not implemented"; return }

func (c *Config) Get(name string) ServerFactory {
	_ = "STUB: not implemented"
	return *new(ServerFactory)
}

func (c *Config) Delete(protocol string) { _ = "STUB: not implemented"; return }

func (c *Config) Load(core Core, protocol string) (server protocol.Server, err error) {
	_ = "STUB: not implemented"
	return *new(protocol.Server), nil
}

func (c *Config) LoadAll(core Core) (serverMap ServerMap, streamServerMap StreamServerMap, err error) {
	_ = "STUB: not implemented"
	return *new(ServerMap), *new(StreamServerMap), nil
}

func New() *Config { _ = "STUB: not implemented"; return nil }
