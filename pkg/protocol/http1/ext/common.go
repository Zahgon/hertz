package ext

import (
	"io"

	errs "github.com/cloudwego/hertz/pkg/common/errors"
	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

const maxContentLengthInStream = 8 * 1024

var errBrokenChunk = errs.NewPublic("cannot find crlf at the end of chunk").SetMeta("when read body chunk")

func MustPeekBuffered(r network.Reader) []byte { _ = "STUB: not implemented"; return nil }

func MustDiscard(r network.Reader, n int) { _ = "STUB: not implemented"; return }

func ReadRawHeaders(dst, buf []byte) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

func WriteBodyChunked(w network.Writer, r io.Reader) error { _ = "STUB: not implemented"; return nil }

func WriteBodyFixedSize(w network.Writer, r io.Reader, size int64) error {
	_ = "STUB: not implemented"
	return nil
}

func appendBodyFixedSize(r network.Reader, dst []byte, n int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func readBodyIdentity(r network.Reader, maxBodySize int, dst []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadBody(r network.Reader, contentLength, maxBodySize int, dst []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func LimitedReaderSize(r io.Reader) int64 { _ = "STUB: not implemented"; return 0 }

func readBodyChunked(r network.Reader, maxBodySize int, dst []byte) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func round2(n int) int { _ = "STUB: not implemented"; return 0 }

func WriteChunk(w network.Writer, b []byte, withFlush bool) error {
	_ = "STUB: not implemented"
	return nil
}

func isOnlyCRLF(b []byte) bool { _ = "STUB: not implemented"; return false }

func BufferSnippet(b []byte) string { _ = "STUB: not implemented"; return "" }

func normalizeHeaderValue(ov, ob []byte, headerLength int) (nv, nb []byte, nhl int) {
	_ = "STUB: not implemented"
	return nil, nil, 0
}

func stripSpace(b []byte) []byte { _ = "STUB: not implemented"; return nil }

func SkipTrailer(r network.Reader) error { _ = "STUB: not implemented"; return nil }

func trySkipTrailer(r network.Reader, n int) error { _ = "STUB: not implemented"; return nil }

func skipTrailer(buf []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func ReadTrailer(t *protocol.Trailer, r network.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func tryReadTrailer(t *protocol.Trailer, r network.Reader, n int) error {
	_ = "STUB: not implemented"
	return nil
}

func parseTrailer(t *protocol.Trailer, buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func WriteTrailer(t *protocol.Trailer, w network.Writer) error {
	_ = "STUB: not implemented"
	return nil
}
