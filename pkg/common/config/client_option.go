package config

import (
	"crypto/tls"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client/retry"
	"github.com/cloudwego/hertz/pkg/network"
)

type ConnPoolState struct {
	PoolConnNum int

	TotalConnNum int

	WaitConnNum int

	Addr string

	MaxConns int
}

type HostClientState interface {
	ConnPoolState() ConnPoolState
}

type HostClientStateFunc func(HostClientState)

type ClientOption struct {
	F func(o *ClientOptions)
}

type ClientOptions struct {
	DialTimeout time.Duration

	MaxConnsPerHost int

	MaxIdleConnDuration time.Duration
	MaxConnDuration     time.Duration
	MaxConnWaitTimeout  time.Duration
	KeepAlive           bool
	ReadTimeout         time.Duration
	TLSConfig           *tls.Config
	ResponseBodyStream  bool

	Name string

	NoDefaultUserAgentHeader bool

	Dialer network.Dialer

	DialDualStack bool

	WriteTimeout time.Duration

	MaxResponseBodySize int

	DisableHeaderNamesNormalizing bool

	DisablePathNormalizing bool

	RetryConfig *retry.Config

	HostClientStateObserve HostClientStateFunc

	ObservationInterval time.Duration

	HostClientConfigHook func(hc interface{}) error
}

func NewClientOptions(opts []ClientOption) *ClientOptions { _ = "STUB: not implemented"; return nil }

func (o *ClientOptions) Apply(opts []ClientOption) { _ = "STUB: not implemented"; return }
