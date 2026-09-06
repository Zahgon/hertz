package ext

import (
	"bytes"
	"io"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

var (
	errChunkedStream = errs.New(errs.ErrChunkedStream, errs.ErrorTypePublic, nil)

	bodyStreamPool = sync.Pool{
		New: func() interface{} {
			return &bodyStream{}
		},
	}
)

var NoBody = protocol.NoBody

type bodyStream struct {
	prefetchedBytes *bytes.Reader
	reader          network.Reader
	trailer         *protocol.Trailer
	offset          int
	contentLength   int
	chunkLeft       int

	chunkEOF bool
}

func ReadBodyWithStreaming(zr network.Reader, contentLength, maxBodySize int, dst []byte) (b []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func AcquireBodyStream(b *bytebufferpool.ByteBuffer, r network.Reader, t *protocol.Trailer, contentLength int) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func (rs *bodyStream) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:errcheck

func (rs *bodyStream) skipRest() error { _ = "STUB: not implemented"; return nil }

func ReleaseBodyStream(requestReader io.Reader) (err error) { _ = "STUB: not implemented"; return nil }

func (rs *bodyStream) reset() { _ = "STUB: not implemented"; return }
