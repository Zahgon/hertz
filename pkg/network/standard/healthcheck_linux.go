//go:build linux

package standard

import (
	"net"
	"sync"
	"syscall"
)

type tcpHealthChecker struct {
	mu      sync.Mutex
	tcpConn *net.TCPConn
	rawConn syscall.RawConn
	healthy bool
	probe   func(uintptr) bool
}

type unhealthyTCPHealthChecker struct{}

func (unhealthyTCPHealthChecker) isHealthy() bool { _ = "STUB: not implemented"; return false }

func newTCPHealthChecker(conn net.Conn) connHealthChecker {
	_ = "STUB: not implemented"
	return *new(connHealthChecker)
}

func (c *tcpHealthChecker) isHealthy() bool { _ = "STUB: not implemented"; return false }

func recvHealthyTCP(fd uintptr) bool { _ = "STUB: not implemented"; return false }
