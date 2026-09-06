//go:build linux

package netpoll

import (
	"time"
)

type reusableConnHealthChecker interface {
	IsHealthyForReuse(ownerWaitTimeout time.Duration) bool
}

func (c *Conn) IsHealthy(probeTimeout, ownerWaitTimeout time.Duration) bool {
	_ = "STUB: not implemented"
	return false
}
