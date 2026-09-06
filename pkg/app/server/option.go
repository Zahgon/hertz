package server

import (
	"context"
	"crypto/tls"
	"net"
	"time"

	"github.com/cloudwego/hertz/pkg/app/server/binding"
	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/tracer"
	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
	"github.com/cloudwego/hertz/pkg/network"
)

func WithKeepAliveTimeout(t time.Duration) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithReadTimeout(t time.Duration) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithWriteTimeout(t time.Duration) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithIdleTimeout(t time.Duration) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithRedirectTrailingSlash(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithRedirectFixedPath(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithHandleMethodNotAllowed(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithUseRawPath(b bool) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithRemoveExtraSlash(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithUnescapePathValues(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithDisablePreParseMultipartForm(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithHostPorts(hp string) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithListener(ln net.Listener) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithBasePath(basePath string) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithMaxRequestBodySize(bs int) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithMaxHeaderBytes(size int) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithMaxKeepBodySize(bs int) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithGetOnly(isOnly bool) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithKeepAlive(b bool) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithStreamBody(b bool) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithNetwork(nw string) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithExitWaitTime(timeout time.Duration) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithTLS(cfg *tls.Config) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithListenConfig(l *net.ListenConfig) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithTransport(transporter func(options *config.Options) network.Transporter) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithAltTransport(transporter func(options *config.Options) network.Transporter) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithH2C(enable bool) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithReadBufferSize(size int) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithALPN(enable bool) config.Option { _ = "STUB: not implemented"; return *new(config.Option) }

func WithTracer(t tracer.Tracer) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithTraceLevel(level stats.Level) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithRegistry(r registry.Registry, info *registry.Info) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithAutoReloadRender(b bool, interval time.Duration) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithDisablePrintRoute(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithOnAccept(fn func(conn net.Conn) context.Context) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithOnConnect(fn func(ctx context.Context, conn network.Conn) context.Context) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithBindConfig(bc *binding.BindConfig) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithValidateConfig(vc *binding.ValidateConfig) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithCustomBinder(b binding.Binder) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithCustomValidator(b binding.StructValidator) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithCustomValidatorFunc(vf binding.ValidatorFunc) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithDisableHeaderNamesNormalizing(disable bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithDisableDefaultDate(disable bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithDisableDefaultContentType(disable bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}

func WithSenseClientDisconnection(b bool) config.Option {
	_ = "STUB: not implemented"
	return *new(config.Option)
}
