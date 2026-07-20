package network

import (
	"io"
	"sync"
)

const size4K = 1024 * 4

type node struct {
	data     []byte
	readOnly bool
}

var nodePool = sync.Pool{}

func init() {
	nodePool.New = func() interface{} {
		return &node{}
	}
}

type networkWriter struct {
	caches []*node
	w      io.Writer
}

func (w *networkWriter) release() { _ = "STUB: not implemented"; return }

func (w *networkWriter) Malloc(length int) (buf []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (w *networkWriter) WriteBinary(b []byte) (length int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (w *networkWriter) Flush() (err error) { _ = "STUB: not implemented"; return nil }

func NewWriter(w io.Writer) Writer { _ = "STUB: not implemented"; return *new(Writer) }

type ExtWriter interface {
	io.Writer

	Flush() error

	Finalize() error
}
