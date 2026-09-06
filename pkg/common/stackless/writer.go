package stackless

import (
	"io"

	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/cloudwego/hertz/pkg/common/errors"
)

type Writer interface {
	Write(p []byte) (int, error)
	Flush() error
	Close() error
	Reset(w io.Writer)
}

type NewWriterFunc func(w io.Writer) Writer

func NewWriter(dstW io.Writer, newWriter NewWriterFunc) Writer {
	_ = "STUB: not implemented"
	return *new(Writer)
}

type writer struct {
	dstW io.Writer
	zw   Writer
	xw   xWriter

	err error
	n   int

	p  []byte
	op op
}

type op int

const (
	opWrite op = iota
	opFlush
	opClose
	opReset
)

func (w *writer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *writer) Flush() error { _ = "STUB: not implemented"; return nil }

func (w *writer) Close() error { _ = "STUB: not implemented"; return nil }

func (w *writer) Reset(dstW io.Writer) { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (w *writer) do(op op) error { _ = "STUB: not implemented"; return nil }

var errHighLoad = errors.NewPublic("cannot compress data due to high load")

var stacklessWriterFunc = NewFunc(writerFunc)

func writerFunc(ctx interface{}) { _ = "STUB: not implemented"; return }

type xWriter struct {
	bb *bytebufferpool.ByteBuffer
}

func (w *xWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w *xWriter) Reset() { _ = "STUB: not implemented"; return }

var bufferPool bytebufferpool.Pool
