package dialer

import (
	"crypto/tls"
	"net"
	"time"

	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/network/standard"
)

var defaultDialer network.Dialer = standard.NewDialer()

func SetDialer(dialer network.Dialer) { _ = "STUB: not implemented"; return }

func DefaultDialer() network.Dialer { _ = "STUB: not implemented"; return *new(network.Dialer) }

func DialConnection(network, address string, timeout time.Duration, tlsConfig *tls.Config) (conn network.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func DialTimeout(network, address string, timeout time.Duration, tlsConfig *tls.Config) (conn net.Conn, err error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil
}

func AddTLS(conn network.Conn, tlsConfig *tls.Config) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}
