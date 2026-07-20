package errors

import (
	"errors"
)

var (
	ErrNeedMore           = errors.New("need more data")
	ErrChunkedStream      = errors.New("chunked stream")
	ErrBodyTooLarge       = errors.New("body size exceeds the given limit")
	ErrHeaderTooLarge     = errors.New("header size exceeds the given limit")
	ErrHijacked           = errors.New("connection has been hijacked")
	ErrTimeout            = errors.New("timeout")
	ErrIdleTimeout        = errors.New("idle timeout")
	ErrNothingRead        = errors.New("nothing read")
	ErrShortConnection    = errors.New("short connection")
	ErrConnectionClosed   = errors.New("connection closed")
	ErrNotSupportProtocol = errors.New("not support protocol")
	ErrNoMultipartForm    = errors.New("request has no multipart/form-data Content-Type")
	ErrBadPoolConn        = errors.New("connection is closed by peer while being in the connection pool")

	ErrNoFreeConns = errors.New("no free connections available to host")
)

type ErrorType uint64

type Error struct {
	Err  error
	Type ErrorType
	Meta interface{}
}

const (
	ErrorTypeBind ErrorType = 1 << iota

	ErrorTypeRender

	ErrorTypePrivate

	ErrorTypePublic

	ErrorTypeAny
)

type ErrorChain []*Error

var _ error = (*Error)(nil)

func (msg *Error) SetType(flags ErrorType) *Error { _ = "STUB: not implemented"; return nil }

func (msg *Error) Error() string { _ = "STUB: not implemented"; return "" }

func (a ErrorChain) String() string { _ = "STUB: not implemented"; return "" }

func (msg *Error) Unwrap() error { _ = "STUB: not implemented"; return nil }

func (msg *Error) SetMeta(data interface{}) *Error { _ = "STUB: not implemented"; return nil }

func (msg *Error) IsType(flags ErrorType) bool { _ = "STUB: not implemented"; return false }

func (msg *Error) JSON() interface{} { _ = "STUB: not implemented"; return nil }

func (a ErrorChain) Errors() []string { _ = "STUB: not implemented"; return nil }

func (a ErrorChain) ByType(typ ErrorType) ErrorChain {
	_ = "STUB: not implemented"
	return *new(ErrorChain)
}

func (a ErrorChain) Last() *Error { _ = "STUB: not implemented"; return nil }

func (a ErrorChain) JSON() interface{} { _ = "STUB: not implemented"; return nil }

func New(err error, t ErrorType, meta interface{}) *Error { _ = "STUB: not implemented"; return nil }

func NewPublic(err string) *Error { _ = "STUB: not implemented"; return nil }

func NewPrivate(err string) *Error { _ = "STUB: not implemented"; return nil }

func Newf(t ErrorType, meta interface{}, format string, v ...interface{}) *Error {
	_ = "STUB: not implemented"
	return nil
}

func NewPublicf(format string, v ...interface{}) *Error { _ = "STUB: not implemented"; return nil }

func NewPrivatef(format string, v ...interface{}) *Error { _ = "STUB: not implemented"; return nil }
