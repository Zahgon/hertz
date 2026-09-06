//go:build !linux

package standard

import "net"

func newTCPHealthChecker(net.Conn) connHealthChecker {
	_ = "STUB: not implemented"
	return *new(connHealthChecker)
}
