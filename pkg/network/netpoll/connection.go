package netpoll

import (
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/netpoll"
)

type Conn struct {
	network.Conn
}

func (c *Conn) ToHertzError(err error) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Peek(n int) (b []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Skip(n int) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Release() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *Conn) ReadByte() (b byte, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) ReadBinary(n int) (b []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) Malloc(n int) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) WriteBinary(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Flush() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) HandleSpecificError(err error, rip string) (needIgnore bool) {
	_ = "STUB: not implemented"
	return false
}

func normalizeErr(err error) error { _ = "STUB: not implemented"; return nil }

func newConn(c netpoll.Connection) network.Conn {
	_ = "STUB: not implemented"
	return *new(network.Conn)
}
