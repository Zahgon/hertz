package mock

import (
	"bufio"
	"bytes"
	"io"
)

type ZeroCopyReader struct {
	*bufio.Reader
}

func (m ZeroCopyReader) Peek(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (m ZeroCopyReader) Skip(n int) (err error) { _ = "STUB: not implemented"; return nil }

func (m ZeroCopyReader) Release() (err error) { _ = "STUB: not implemented"; return nil }

func (m ZeroCopyReader) Len() (length int) { _ = "STUB: not implemented"; return 0 }

func (m ZeroCopyReader) ReadBinary(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewZeroCopyReader(r string) ZeroCopyReader {
	_ = "STUB: not implemented"
	return *new(ZeroCopyReader)
}

func NewLimitReader(r *bytes.Buffer) io.LimitedReader {
	_ = "STUB: not implemented"
	return *new(io.LimitedReader)
}

type EOFReader struct{}

func (e *EOFReader) Peek(n int) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (e *EOFReader) Skip(n int) error { _ = "STUB: not implemented"; return nil }

func (e *EOFReader) Release() error { _ = "STUB: not implemented"; return nil }

func (e *EOFReader) Len() int { _ = "STUB: not implemented"; return 0 }

func (e *EOFReader) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (e *EOFReader) ReadBinary(n int) (p []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e *EOFReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }
