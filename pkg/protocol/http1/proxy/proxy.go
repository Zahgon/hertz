package proxy

import (
	"crypto/tls"

	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

func SetupProxy(conn network.Conn, addr string, proxyURI *protocol.URI, tlsConfig *tls.Config, isTLS bool, dialer network.Dialer) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func SetProxyAuthHeader(h *protocol.RequestHeader, proxyURI *protocol.URI) {
	_ = "STUB: not implemented"
	return
}
