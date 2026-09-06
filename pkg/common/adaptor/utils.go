package adaptor

import (
	"bytes"
	"io"

	"github.com/cloudwego/hertz/pkg/network"
)

func methodstr(m []byte) string { _ = "STUB: not implemented"; return "" }

type bytesRWCloser struct {
	bytes.Reader
}

func newBytesRWCloser(b []byte) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func (bytesRWCloser) Close() error { _ = "STUB: not implemented"; return nil }

func writer2writerExt(w network.Writer) network.ExtWriter {
	_ = "STUB: not implemented"
	return *new(network.ExtWriter)
}

type extWriter struct {
	network.Writer
}

func (w extWriter) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (w extWriter) Finalize() error { _ = "STUB: not implemented"; return nil }

func reader2closer(r io.Reader) io.ReadCloser {
	_ = "STUB: not implemented"
	return *new(io.ReadCloser)
}

func parseHTTPVersion(s string) (major, minor int, _ error) {
	_ = "STUB: not implemented"
	return 0, 0, nil
}
