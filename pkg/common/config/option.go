package config

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/network"
)

type Option struct {
	F func(o *Options)
}

const (
	defaultKeepAliveTimeout   = 1 * time.Minute
	defaultReadTimeout        = 3 * time.Minute
	defaultAddr               = ":8888"
	defaultNetwork            = "tcp"
	defaultBasePath           = "/"
	defaultMaxRequestBodySize = 4 << 20
	defaultMaxHeaderBytes     = 1 << 20
	defaultWaitExitTimeout    = time.Second * 5
	defaultReadBufferSize     = 4 * 1024
)

type Options struct {
	KeepAliveTimeout             time.Duration
	ReadTimeout                  time.Duration
	WriteTimeout                 time.Duration
	IdleTimeout                  time.Duration
	RedirectTrailingSlash        bool
	MaxRequestBodySize           int
	MaxHeaderBytes               int
	MaxKeepBodySize              int
	GetOnly                      bool
	DisableKeepalive             bool
	RedirectFixedPath            bool
	HandleMethodNotAllowed       bool
	UseRawPath                   bool
	RemoveExtraSlash             bool
	UnescapePathValues           bool
	DisablePreParseMultipartForm bool
	NoDefaultDate                bool
	NoDefaultContentType         bool
	StreamRequestBody            bool
	NoDefaultServerHeader        bool
	DisablePrintRoute            bool
	SenseClientDisconnection     bool
	Network                      string
	Addr                         string
	BasePath                     string
	ExitWaitTimeout              time.Duration
	TLS                          *tls.Config
	H2C                          bool
	ReadBufferSize               int
	ALPN                         bool
	Tracers                      []interface{}
	TraceLevel                   interface{}
	Listener                     net.Listener
	ListenConfig                 *net.ListenConfig
	BindConfig                   interface{}
	CustomBinder                 interface{}
	CustomValidator              interface{}

	ValidateConfig interface{}

	TransporterNewer    func(opt *Options) network.Transporter
	AltTransporterNewer func(opt *Options) network.Transporter

	OnAccept  func(conn net.Conn) context.Context
	OnConnect func(ctx context.Context, conn network.Conn) context.Context

	Registry registry.Registry

	RegistryInfo *registry.Info

	AutoReloadRender bool

	AutoReloadInterval time.Duration

	DisableHeaderNamesNormalizing bool
}

func (o *Options) Apply(opts []Option) { _ = "STUB: not implemented"; return }

func NewOptions(opts []Option) *Options { _ = "STUB: not implemented"; return nil }
