package sse

import (
	"bufio"
	"context"
	"errors"
	"io"

	"github.com/cloudwego/hertz/pkg/protocol"
)

var errNotSSEContentType = errors.New("Content-Type returned by server is NOT text/event-stream")

type Reader struct {
	resp   *protocol.Response
	r      io.Reader
	s      *bufio.Scanner
	events int32

	lastEventID string
}

func NewReader(resp *protocol.Response) (*Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reader) SetMaxBufferSize(max int) { _ = "STUB: not implemented"; return }

type forceCloseIf interface {
	ForceClose() error
}

func (r *Reader) ForEach(ctx context.Context, f func(e *Event) error) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reader) LastEventID() string { _ = "STUB: not implemented"; return "" }

func (r *Reader) onEventRead(e *Event) { _ = "STUB: not implemented"; return }

func (r *Reader) Read(e *Event) error { _ = "STUB: not implemented"; return nil }

func (r *Reader) Close() error { _ = "STUB: not implemented"; return nil }
