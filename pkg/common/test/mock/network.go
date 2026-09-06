package mock

import (
	"crypto/tls"
	"net"
	"time"

	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/netpoll"
)

var (
	ErrReadTimeout  = errs.New(errs.ErrTimeout, errs.ErrorTypePublic, "read timeout")
	ErrWriteTimeout = errs.New(errs.ErrTimeout, errs.ErrorTypePublic, "write timeout")
)

type Conn struct {
	zr       network.Reader
	zw       network.ReadWriter
	wroteLen int

	rtimeout time.Duration
	wtimeout time.Duration
}

type Recorder interface {
	network.Reader
	WroteLen() int
}

func (m *Conn) SetWriteTimeout(t time.Duration) error { _ = "STUB: not implemented"; return nil }

type SlowReadConn struct {
	*Conn
}

func (m *SlowReadConn) SetWriteTimeout(t time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (m *SlowReadConn) SetReadTimeout(t time.Duration) error { _ = "STUB: not implemented"; return nil }

func SlowReadDialer(addr string) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func SlowWriteDialer(addr string) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (m *Conn) ReadBinary(n int) (p []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Conn) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Conn) Write(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Conn) Release() error { _ = "STUB: not implemented"; return nil }

func (m *Conn) Peek(i int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Conn) Skip(n int) error { _ = "STUB: not implemented"; return nil }

func (m *Conn) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Conn) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *Conn) Malloc(n int) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (m *Conn) WriteBinary(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (m *Conn) Flush() error { _ = "STUB: not implemented"; return nil }

func (m *Conn) WriterRecorder() Recorder { _ = "STUB: not implemented"; return *new(Recorder) }

func (m *Conn) GetReadTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

func (m *Conn) GetWriteTimeout() time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

type recorder struct {
	c *Conn
	network.Reader
}

func (r *recorder) WroteLen() int { _ = "STUB: not implemented"; return 0 }

func (m *SlowReadConn) Peek(i int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewConn(source string) *Conn { _ = "STUB: not implemented"; return nil }

type BrokenConn struct {
	*Conn
}

func (o *BrokenConn) Peek(i int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *BrokenConn) Read(b []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (o *BrokenConn) Flush() error { _ = "STUB: not implemented"; return nil }

func NewBrokenConn(source string) *BrokenConn { _ = "STUB: not implemented"; return nil }

type OneTimeConn struct {
	isRead        bool
	isFlushed     bool
	contentLength int
	*Conn
}

func (o *OneTimeConn) Peek(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (o *OneTimeConn) Skip(n int) error { _ = "STUB: not implemented"; return nil }

func (o *OneTimeConn) Flush() error { _ = "STUB: not implemented"; return nil }

func NewOneTimeConn(source string) *OneTimeConn { _ = "STUB: not implemented"; return nil }

func NewSlowReadConn(source string) *SlowReadConn { _ = "STUB: not implemented"; return nil }

type ErrorReadConn struct {
	*Conn
	errorToReturn error
}

func NewErrorReadConn(err error) *ErrorReadConn { _ = "STUB: not implemented"; return nil }

func (er *ErrorReadConn) Peek(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type SlowWriteConn struct {
	*Conn
	writeTimeout time.Duration
}

func (m *SlowWriteConn) SetWriteTimeout(t time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSlowWriteConn(source string) *SlowWriteConn { _ = "STUB: not implemented"; return nil }

func (m *SlowWriteConn) Flush() error { _ = "STUB: not implemented"; return nil }

func (m *Conn) Close() error { _ = "STUB: not implemented"; return nil }

func (m *Conn) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (m *Conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (m *Conn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (m *Conn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (m *Conn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (m *Conn) Reader() network.Reader { _ = "STUB: not implemented"; return *new(network.Reader) }

func (m *Conn) Writer() network.Writer { _ = "STUB: not implemented"; return *new(network.Writer) }

func (m *Conn) IsActive() bool { _ = "STUB: not implemented"; return false }

func (m *Conn) SetIdleTimeout(timeout time.Duration) error { _ = "STUB: not implemented"; return nil }

func (m *Conn) SetReadTimeout(t time.Duration) error { _ = "STUB: not implemented"; return nil }

func (m *Conn) SetOnRequest(on netpoll.OnRequest) error { _ = "STUB: not implemented"; return nil }

func (m *Conn) AddCloseCallback(callback netpoll.CloseCallback) error {
	_ = "STUB: not implemented"
	return nil
}

type StreamConn struct {
	HasReleased bool
	Data        []byte
}

func NewStreamConn() *StreamConn { _ = "STUB: not implemented"; return nil }

func (m *StreamConn) Peek(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m *StreamConn) Skip(n int) error { _ = "STUB: not implemented"; return nil }

func (m *StreamConn) Release() error { _ = "STUB: not implemented"; return nil }

func (m *StreamConn) Len() int { _ = "STUB: not implemented"; return 0 }

func (m *StreamConn) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (m *StreamConn) ReadBinary(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func DialerFun(addr string) (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

type MockWriter struct {
	w network.Writer

	MockMalloc      func(n int) (buf []byte, err error)
	MockWriteBinary func(b []byte) (n int, err error)
	MockFlush       func() error
}

func NewMockWriter(w network.Writer) *MockWriter { _ = "STUB: not implemented"; return nil }

func (m *MockWriter) Malloc(n int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (m *MockWriter) WriteBinary(b []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (m *MockWriter) Flush() error { _ = "STUB: not implemented"; return nil }

type TLSConn struct {
	network.Conn

	HandshakeErr error
}

var _ network.ConnTLSer = (*TLSConn)(nil)

func (c *TLSConn) Handshake() error { _ = "STUB: not implemented"; return nil }

func (c *TLSConn) ConnectionState() tls.ConnectionState {
	_ = "STUB: not implemented"
	return *new(tls.ConnectionState)
}

func NewTLSConn(conn network.Conn) *TLSConn { _ = "STUB: not implemented"; return nil }
