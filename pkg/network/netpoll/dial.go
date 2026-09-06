package netpoll

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/netpoll"
)

var errNotSupportTLS = errors.NewPublic("not support tls")

type dialer struct {
	netpoll.Dialer
}

func (d dialer) DialConnection(n, address string, timeout time.Duration, tlsConfig *tls.Config) (conn network.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (d dialer) DialTimeout(network, address string, timeout time.Duration, tlsConfig *tls.Config) (conn net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func (d dialer) AddTLS(conn network.Conn, tlsConfig *tls.Config) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func NewDialer() network.Dialer { _ = "STUB: not implemented"; return *new(network.Dialer) }
