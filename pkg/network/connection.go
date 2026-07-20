package network

import (
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"time"
)

type Reader interface {
	Peek(n int) ([]byte, error)

	Skip(n int) error

	Release() error

	Len() int

	ReadByte() (byte, error)

	ReadBinary(n int) (p []byte, err error)
}

type Writer interface {
	Malloc(n int) (buf []byte, err error)

	WriteBinary(b []byte) (n int, err error)

	Flush() error
}

type ReadWriter interface {
	Reader
	Writer
}

type Conn interface {
	net.Conn
	Reader
	Writer

	SetReadTimeout(t time.Duration) error
	SetWriteTimeout(t time.Duration) error
}

type ConnTLSer interface {
	Handshake() error
	ConnectionState() tls.ConnectionState
}

type HandleSpecificError interface {
	HandleSpecificError(err error, rip string) (needIgnore bool)
}

type ErrorNormalization interface {
	ToHertzError(err error) error
}

type DialFunc func(addr string) (Conn, error)

type StreamConn interface {
	GetRawConnection() interface{}

	HandshakeComplete() context.Context

	GetVersion() uint32

	CloseWithError(err ApplicationError, errMsg string) error

	LocalAddr() net.Addr

	RemoteAddr() net.Addr

	Context() context.Context

	Streamer
}

type Streamer interface {
	AcceptStream(context.Context) (Stream, error)

	AcceptUniStream(context.Context) (ReceiveStream, error)

	OpenStream() (Stream, error)

	OpenStreamSync(context.Context) (Stream, error)

	OpenUniStream() (SendStream, error)

	OpenUniStreamSync(context.Context) (SendStream, error)
}

type Stream interface {
	ReceiveStream
	SendStream
}

type ReceiveStream interface {
	StreamID() int64
	io.Reader

	CancelRead(err ApplicationError)

	SetReadDeadline(t time.Time) error
}

type SendStream interface {
	StreamID() int64

	io.Writer

	CancelWrite(err ApplicationError)

	io.Closer

	Context() context.Context

	SetWriteDeadline(t time.Time) error
}

type ApplicationError interface {
	ErrCode() uint64
	fmt.Stringer
}
