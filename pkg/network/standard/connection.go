package standard

import (
	"crypto/tls"
	"io"
	"net"
	"time"

	"github.com/cloudwego/gopkg/bufiox"
	"github.com/cloudwego/gopkg/net/connstate"

	"github.com/cloudwego/hertz/pkg/network"
)

type Conn struct {
	c      net.Conn
	br     *bufiox.DefaultReader
	bw     *bufiox.DefaultWriter
	stater connstate.ConnStater

	buf [8]byte
}

func (c *Conn) ToHertzError(err error) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetWriteTimeout(t time.Duration) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetReadTimeout(t time.Duration) error { _ = "STUB: not implemented"; return nil }

type TLSConn struct {
	Conn
}

func (c *Conn) Peek(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) Skip(n int) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Release() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Len() int { _ = "STUB: not implemented"; return 0 }

func (c *Conn) ReadByte() (b byte, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) ReadBinary(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) Read(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) ReadFrom(r io.Reader) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) Malloc(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) WriteBinary(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *Conn) Flush() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) HandleSpecificError(err error, rip string) (needIgnore bool) {
	_ = "STUB: not implemented"
	return false
}

func (c *TLSConn) Handshake() error { _ = "STUB: not implemented"; return nil }

func (c *TLSConn) ConnectionState() tls.ConnectionState {
	_ = "STUB: not implemented"
	return *new(tls.ConnectionState)
}

func newConn(c net.Conn, size int) network.Conn {
	_ = "STUB: not implemented"
	return *new(network.Conn)
}

func newTLSConn(c net.Conn, size int) network.Conn {
	_ = "STUB: not implemented"
	return *new(network.Conn)
}
