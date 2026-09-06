package client

import (
	"context"
	"sync"
	"time"

	"github.com/cloudwego/hertz/internal/nocopy"
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/client"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/protocol/http1"
	"github.com/cloudwego/hertz/pkg/protocol/suite"
)

var (
	errorInvalidURI          = errors.NewPublic("invalid uri")
	errorLastMiddlewareExist = errors.NewPublic("last middleware already set")
)

func Do(ctx context.Context, req *protocol.Request, resp *protocol.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func DoTimeout(ctx context.Context, req *protocol.Request, resp *protocol.Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func DoDeadline(ctx context.Context, req *protocol.Request, resp *protocol.Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func DoRedirects(ctx context.Context, req *protocol.Request, resp *protocol.Response, maxRedirectsCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func Get(ctx context.Context, dst []byte, url string, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func GetTimeout(ctx context.Context, dst []byte, url string, timeout time.Duration, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func GetDeadline(ctx context.Context, dst []byte, url string, deadline time.Time, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func Post(ctx context.Context, dst []byte, url string, postArgs *protocol.Args, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

var defaultClient, _ = NewClient(WithDialTimeout(consts.DefaultDialTimeout))

type Client struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	options *config.ClientOptions

	Proxy protocol.Proxy

	RetryIfFunc client.RetryIfFunc

	clientFactory suite.ClientFactory

	mLock          sync.Mutex
	m              map[string]client.HostClient
	ms             map[string]client.HostClient
	mws            Middleware
	lastMiddleware Middleware
}

func (c *Client) GetOptions() *config.ClientOptions { _ = "STUB: not implemented"; return nil }

func (c *Client) SetRetryIfFunc(retryIf client.RetryIfFunc) { _ = "STUB: not implemented"; return }

func (c *Client) SetRetryIf(fn func(request *protocol.Request) bool) {
	_ = "STUB: not implemented"
	return
}

func (c *Client) SetProxy(p protocol.Proxy) { _ = "STUB: not implemented"; return }

func (c *Client) Get(ctx context.Context, dst []byte, url string, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) GetTimeout(ctx context.Context, dst []byte, url string, timeout time.Duration, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) GetDeadline(ctx context.Context, dst []byte, url string, deadline time.Time, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) Post(ctx context.Context, dst []byte, url string, postArgs *protocol.Args, requestOptions ...config.RequestOption) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *Client) DoTimeout(ctx context.Context, req *protocol.Request, resp *protocol.Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DoDeadline(ctx context.Context, req *protocol.Request, resp *protocol.Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) DoRedirects(ctx context.Context, req *protocol.Request, resp *protocol.Response, maxRedirectsCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) Do(ctx context.Context, req *protocol.Request, resp *protocol.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) do(ctx context.Context, req *protocol.Request, resp *protocol.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Client) CloseIdleConnections() { _ = "STUB: not implemented"; return }

func (c *Client) cleaner(isTLS bool) { _ = "STUB: not implemented"; return }

func (c *Client) cleanHostClients(isTLS bool) bool { _ = "STUB: not implemented"; return false }

func (c *Client) SetClientFactory(cf suite.ClientFactory) { _ = "STUB: not implemented"; return }

func (c *Client) GetDialerName() (dName string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func NewClient(opts ...config.ClientOption) (*Client, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Client) Use(mws ...Middleware) { _ = "STUB: not implemented"; return }

func (c *Client) UseAsLast(mw Middleware) error { _ = "STUB: not implemented"; return nil }

func (c *Client) TakeOutLastMiddleware() Middleware {
	_ = "STUB: not implemented"
	return *new(Middleware)
}

func newHttp1OptionFromClient(c *Client) *http1.ClientOptions {
	_ = "STUB: not implemented"
	return nil
}
