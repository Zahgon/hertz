package compress

import (
	"compress/gzip"
	"io"
	"sync"

	"github.com/cloudwego/hertz/pkg/common/stackless"
)

const CompressDefaultCompression = 6

var gzipReaderPool sync.Pool

var (
	stacklessGzipWriterPoolMap = newCompressWriterPoolMap()
	realGzipWriterPoolMap      = newCompressWriterPoolMap()
)

func newCompressWriterPoolMap() []*sync.Pool { _ = "STUB: not implemented"; return nil }

type compressCtx struct {
	w     io.Writer
	p     []byte
	level int
}

func AppendGunzipBytes(dst, src []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

type byteSliceWriter struct {
	b []byte
}

func (w *byteSliceWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func WriteGunzip(w io.Writer, p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type byteSliceReader struct {
	b []byte
}

func (r *byteSliceReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func AcquireGzipReader(r io.Reader) (*gzip.Reader, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReleaseGzipReader(zr *gzip.Reader) { _ = "STUB: not implemented"; return }

func AppendGzipBytes(dst, src []byte) []byte { _ = "STUB: not implemented"; return nil }

func AppendGzipBytesLevel(dst, src []byte, level int) []byte { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

var stacklessWriteGzip = stackless.NewFunc(nonblockingWriteGzip)

func nonblockingWriteGzip(ctxv interface{}) { _ = "STUB: not implemented"; return }

func releaseRealGzipWriter(zw *gzip.Writer, level int) { _ = "STUB: not implemented"; return }

func acquireRealGzipWriter(w io.Writer, level int) *gzip.Writer {
	_ = "STUB: not implemented"
	return nil
}

func normalizeCompressLevel(level int) int { _ = "STUB: not implemented"; return 0 }

func WriteGzipLevel(w io.Writer, p []byte, level int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func AcquireStacklessGzipWriter(w io.Writer, level int) stackless.Writer {
	_ = "STUB: not implemented"
	return *new(stackless.Writer)
}

func ReleaseStacklessGzipWriter(sw stackless.Writer, level int) { _ = "STUB: not implemented"; return }
