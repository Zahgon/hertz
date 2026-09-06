package http1

import (
	"context"
	"crypto/tls"
	"sync"
	"sync/atomic"
	"time"

	"github.com/cloudwego/hertz/internal/nocopy"
	"github.com/cloudwego/hertz/pkg/app/client/retry"
	"github.com/cloudwego/hertz/pkg/common/config"
	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/client"
)

var (
	errConnectionClosed = errs.NewPublic("the server closed connection before returning the first response byte. " +
		"Make sure the server returns 'Connection: close' response header before closing the connection")

	errTimeout = errs.New(errs.ErrTimeout, errs.ErrorTypePublic, "host client")
)

type HostClient struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	*ClientOptions

	Addr     string
	IsTLS    bool
	ProxyURI *protocol.URI

	clientName  atomic.Value
	lastUseTime uint32

	connsLock  sync.Mutex
	connsCount int
	conns      []*clientConn
	connsWait  *wantConnQueue

	addrsLock sync.Mutex
	addrs     []string
	addrIdx   uint32

	tlsConfigMap     map[string]*tls.Config
	tlsConfigMapLock sync.Mutex

	pendingRequests int32

	connsCleanerRun bool

	closed chan struct{}
}

func (c *HostClient) SetDynamicConfig(dc *client.DynamicConfig) { _ = "STUB: not implemented"; return }

type clientConn struct {
	c network.Conn

	createdTime time.Time
	lastUseTime time.Time
}

var startTimeUnix = time.Now().Unix()

func (c *HostClient) LastUseTime() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (c *HostClient) Get(ctx context.Context, dst []byte, url string) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *HostClient) ConnectionCount() (count int) { _ = "STUB: not implemented"; return 0 }

func (c *HostClient) WantConnectionCount() (count int) { _ = "STUB: not implemented"; return 0 }

func (c *HostClient) ConnPoolState() config.ConnPoolState {
	_ = "STUB: not implemented"
	return *new(config.ConnPoolState)
}

func (c *HostClient) GetTimeout(ctx context.Context, dst []byte, url string, timeout time.Duration) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *HostClient) GetDeadline(ctx context.Context, dst []byte, url string, deadline time.Time) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

func (c *HostClient) Post(ctx context.Context, dst []byte, url string, postArgs *protocol.Args) (statusCode int, body []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

type wantConnQueue struct {
	head    []*wantConn
	headPos int
	tail    []*wantConn
}

type wantConn struct {
	ready chan struct{}
	mu    sync.Mutex
	conn  *clientConn
	err   error
}

func (c *HostClient) DoTimeout(ctx context.Context, req *protocol.Request, resp *protocol.Response, timeout time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HostClient) DoDeadline(ctx context.Context, req *protocol.Request, resp *protocol.Response, deadline time.Time) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HostClient) DoRedirects(ctx context.Context, req *protocol.Request, resp *protocol.Response, maxRedirectsCount int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *HostClient) Do(ctx context.Context, req *protocol.Request, resp *protocol.Response) error {
	_ = "STUB: not implemented"
	return nil
}

//nolint:errcheck

func (c *HostClient) PendingRequests() int { _ = "STUB: not implemented"; return 0 }

func (c *HostClient) do(req *protocol.Request, resp *protocol.Response) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func timeUntil(deadline time.Time) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func calcTimeout(deadline time.Time, timeout time.Duration) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (c *HostClient) getTimeouts(o *config.RequestOptions) (dtimeout, rtimeout, wtimeout time.Duration) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(time.Duration), *new(time.Duration)
}

func (c *HostClient) doNonNilReqResp(req *protocol.Request, resp *protocol.Response) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

//nolint:errcheck

//nolint:errcheck

var poolUpgradeConn = sync.Pool{
	New: func() interface{} {
		return &upgradeConn{}
	},
}

type upgradeConn struct {
	c  *HostClient
	cc *clientConn
}

func newUpgradeConn(c *HostClient, cc *clientConn) *upgradeConn {
	_ = "STUB: not implemented"
	return nil
}

func (p *upgradeConn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (p *upgradeConn) Hijack() (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (p *upgradeConn) gc() error { _ = "STUB: not implemented"; return nil }

func (c *HostClient) Close() error { _ = "STUB: not implemented"; return nil }

func (c *HostClient) SetMaxConns(newMaxConns int) { _ = "STUB: not implemented"; return }

const pooledConnHealthCheckTimeout = 50 * time.Microsecond

type pooledConnHealthChecker interface {
	IsHealthy(probeTimeout, ownerWaitTimeout time.Duration) bool
}

func (c *HostClient) acquireConn(dialTimeout time.Duration, deadline time.Time) (cc *clientConn, inPool bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *HostClient) acquireConnOnce(dialTimeout time.Duration) (cc *clientConn, inPool bool, err error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}

func (c *HostClient) isPooledConnHealthy(conn network.Conn, probeTimeout, ownerWaitTimeout time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *HostClient) queueForIdle(w *wantConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) dialConnFor(w *wantConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) CloseIdleConnections() { _ = "STUB: not implemented"; return }

func (c *HostClient) ShouldRemove() bool { _ = "STUB: not implemented"; return false }

func (c *HostClient) connsCleaner() { _ = "STUB: not implemented"; return }

func (c *HostClient) closeConn(cc *clientConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) decConnsCount() { _ = "STUB: not implemented"; return }

func acquireClientConn(conn network.Conn) *clientConn { _ = "STUB: not implemented"; return nil }

func releaseClientConn(cc *clientConn) { _ = "STUB: not implemented"; return }

var clientConnPool sync.Pool

func (c *HostClient) releaseConn(cc *clientConn) { _ = "STUB: not implemented"; return }

func (c *HostClient) acquireWriter(conn network.Conn) network.Writer {
	_ = "STUB: not implemented"
	return *new(network.Writer)
}

func (c *HostClient) acquireReader(conn network.Conn) network.Reader {
	_ = "STUB: not implemented"
	return *new(network.Reader)
}

func newClientTLSConfig(c *tls.Config, addr string) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

func tlsServerName(addr string) string { _ = "STUB: not implemented"; return "" }

func (c *HostClient) nextAddr() string { _ = "STUB: not implemented"; return "" }

func (c *HostClient) dialHostHard(dialTimeout time.Duration) (conn network.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (c *HostClient) cachedTLSConfig(addr string) *tls.Config {
	_ = "STUB: not implemented"
	return nil
}

func dialAddr(addr string, dial network.Dialer, dialDualStack bool, tlsConfig *tls.Config, timeout time.Duration, proxyURI *protocol.URI, isTLS bool) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (c *HostClient) getClientName() []byte { _ = "STUB: not implemented"; return nil }

func (w *wantConn) waiting() bool { _ = "STUB: not implemented"; return false }

func (w *wantConn) tryDeliver(conn *clientConn, err error) bool {
	_ = "STUB: not implemented"
	return false
}

func (w *wantConn) cancel(c *HostClient, err error) { _ = "STUB: not implemented"; return }

func (q *wantConnQueue) len() int { _ = "STUB: not implemented"; return 0 }

func (q *wantConnQueue) pushBack(w *wantConn) { _ = "STUB: not implemented"; return }

func (q *wantConnQueue) popFront() *wantConn { _ = "STUB: not implemented"; return nil }

func (q *wantConnQueue) peekFront() *wantConn { _ = "STUB: not implemented"; return nil }

func (q *wantConnQueue) clearFront() (cleaned bool) { _ = "STUB: not implemented"; return false }

func NewHostClient(c *ClientOptions) client.HostClient {
	_ = "STUB: not implemented"
	return *new(client.HostClient)
}

type ClientOptions struct {
	Name string

	NoDefaultUserAgentHeader bool

	Dialer network.Dialer

	DialTimeout time.Duration

	DialDualStack bool

	TLSConfig *tls.Config

	MaxConns int

	MaxConnDuration time.Duration

	MaxIdleConnDuration time.Duration

	PooledConnHealthCheck bool

	ReadTimeout time.Duration

	WriteTimeout time.Duration

	MaxResponseBodySize int

	DisableHeaderNamesNormalizing bool

	DisablePathNormalizing bool

	MaxConnWaitTimeout time.Duration

	ResponseBodyStream bool

	RetryConfig *retry.Config

	RetryIfFunc client.RetryIfFunc

	StateObserve config.HostClientStateFunc

	ObservationInterval time.Duration
}
