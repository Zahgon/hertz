package client

import (
	"crypto/tls"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client/retry"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/network"
)

func WithDialTimeout(dialTimeout time.Duration) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithMaxConnsPerHost(mc int) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithMaxIdleConnDuration(t time.Duration) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithMaxConnDuration(t time.Duration) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithMaxConnWaitTimeout(t time.Duration) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithKeepAlive(b bool) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithClientReadTimeout(t time.Duration) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithTLSConfig(cfg *tls.Config) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithDialer(d network.Dialer) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithResponseBodyStream(b bool) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithHostClientConfigHook(h func(hc interface{}) error) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithDisableHeaderNamesNormalizing(disable bool) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithName(name string) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithNoDefaultUserAgentHeader(isNoDefaultUserAgentHeader bool) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithDisablePathNormalizing(isDisablePathNormalizing bool) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithRetryConfig(opts ...retry.Option) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithWriteTimeout(t time.Duration) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithConnStateObserve(hs config.HostClientStateFunc, interval ...time.Duration) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

func WithDialFunc(f network.DialFunc, dialers ...network.Dialer) config.ClientOption {
	_ = "STUB: not implemented"
	return *new(config.ClientOption)
}

type customDialer struct {
	network.Dialer
	dialFunc network.DialFunc
}

func (m *customDialer) DialConnection(network, address string, timeout time.Duration, tlsConfig *tls.Config) (conn network.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func newCustomDialerWithDialFunc(dialer network.Dialer, dialFunc network.DialFunc) network.Dialer {
	_ = "STUB: not implemented"
	return *new(network.Dialer)
}
