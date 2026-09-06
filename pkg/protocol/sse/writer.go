package sse

import (
	"errors"
	"sync"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/network"
)

type Writer struct {
	w network.ExtWriter

	mu sync.Mutex
}

func NewWriter(c *app.RequestContext) *Writer { _ = "STUB: not implemented"; return nil }

var (
	errIDContainsCRLR   = errors.New(`id field contains '\r' or '\n'`)
	errTypeContainsCRLR = errors.New(`event field contains '\r' or '\n'`)
)

func (w *Writer) WriteEvent(id, eventType string, data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (w *Writer) WriteKeepAlive() error { _ = "STUB: not implemented"; return nil }

func (w *Writer) WriteComment(s string) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) Write(e *Event) error { _ = "STUB: not implemented"; return nil }

func (w *Writer) Close() error { _ = "STUB: not implemented"; return nil }
