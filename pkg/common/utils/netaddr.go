package utils

import "net"

var _ net.Addr = (*NetAddr)(nil)

type NetAddr struct {
	network string
	address string
}

func NewNetAddr(network, address string) net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (na *NetAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (na *NetAddr) String() string { _ = "STUB: not implemented"; return "" }
