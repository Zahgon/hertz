package client

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol"
)

const defaultMaxRedirectsCount = 16

var (
	errTimeout          = errors.New(errors.ErrTimeout, errors.ErrorTypePublic, "host client")
	errMissingLocation  = errors.NewPublic("missing Location header for http redirect")
	errTooManyRedirects = errors.NewPublic("too many redirects detected when doing the request")

	clientURLResponseChPool sync.Pool
)

type HostClient interface {
	Doer
	SetDynamicConfig(dc *DynamicConfig)
	CloseIdleConnections()
	ShouldRemove() bool
	ConnectionCount() int
}

type Doer interface {
	Do(ctx context.Context, req *protocol.Request, resp *protocol.Response) error
}

func DefaultRetryIf(req *protocol.Request, resp *protocol.Response, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func isIdempotent(req *protocol.Request, resp *protocol.Response, err error) bool {
	_ = "STUB: not implemented"
	return false
}

type DynamicConfig struct {
	Addr     string
	ProxyURI *protocol.URI
	IsTLS    bool
}

type RetryIfFunc func(req *protocol.Request, resp *protocol.Response, err error) bool

type clientURLResponse struct {
	statusCode int
	body       []byte
	err        error
}

func GetURL(ctx context.Context, dst []byte, url string, c Doer, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func GetURLTimeout(ctx context.Context, dst []byte, url string, timeout time.Duration, c Doer, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func GetURLDeadline(ctx context.Context, dst []byte, url string, deadline time.Time, c Doer, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func PostURL(ctx context.Context, dst []byte, url string, postArgs *protocol.Args, c Doer, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func doRequestFollowRedirectsBuffer(ctx context.Context, req *protocol.Request, dst []byte, url string, c Doer) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func DoRequestFollowRedirects(ctx context.Context, req *protocol.Request, resp *protocol.Response, url string, maxRedirectsCount int, c Doer) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func StatusCodeIsRedirect(statusCode int) bool { _ = "STUB: not implemented"; return false }

func getRedirectURL(baseURL string, location []byte) string { _ = "STUB: not implemented"; return "" }

func DoTimeout(ctx context.Context, req *protocol.Request, resp *protocol.Response, timeout time.Duration, c Doer) error {
	_ = "STUB: not implemented"
	return nil
}

func DoDeadline(ctx context.Context, req *protocol.Request, resp *protocol.Response, deadline time.Time, c Doer) error {
	_ = "STUB: not implemented"
	return nil
}
