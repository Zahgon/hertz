package network

import (
	"crypto/tls"
	"net"
	"time"
)

type Dialer interface {
	DialConnection(network, address string, timeout time.Duration, tlsConfig *tls.Config) (conn Conn, err error)

	DialTimeout(network, address string, timeout time.Duration, tlsConfig *tls.Config) (conn net.Conn, err error)

	AddTLS(conn Conn, tlsConfig *tls.Config) (Conn, error)
}
