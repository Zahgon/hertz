package resp

import (
	"errors"
	"sync"

	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var chunkWriterPool sync.Pool

func init() {
	chunkWriterPool = sync.Pool{
		New: func() interface{} {
			return &chunkedBodyWriter{}
		},
	}
}

type chunkedBodyWriter struct {
	r *protocol.Response
	w network.Writer

	err         error
	finalized   bool
	wroteHeader bool
}

var errChunkedFinished = errors.New("chunked response is finished; no more data will be written.")

func (c *chunkedBodyWriter) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *chunkedBodyWriter) WriteHeader() error { _ = "STUB: not implemented"; return nil }

func (c *chunkedBodyWriter) writeChunk(b []byte) error { _ = "STUB: not implemented"; return nil }

func (c *chunkedBodyWriter) Flush() error { _ = "STUB: not implemented"; return nil }

func (c *chunkedBodyWriter) Finalize() error { _ = "STUB: not implemented"; return nil }

func (c *chunkedBodyWriter) release() { _ = "STUB: not implemented"; return }

func NewChunkedBodyWriter(r *protocol.Response, w network.Writer) network.ExtWriter {
	_ = "STUB: not implemented"
	return *new(network.ExtWriter)
}
