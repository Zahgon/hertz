package bytebufferpool

import "io"

type ByteBuffer struct {
	B []byte
}

func (b *ByteBuffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *ByteBuffer) Cap() int { _ = "STUB: not implemented"; return 0 }

func (b *ByteBuffer) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *ByteBuffer) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *ByteBuffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *ByteBuffer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *ByteBuffer) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

func (b *ByteBuffer) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *ByteBuffer) Set(p []byte) { _ = "STUB: not implemented"; return }

func (b *ByteBuffer) SetString(s string) { _ = "STUB: not implemented"; return }

func (b *ByteBuffer) String() string { _ = "STUB: not implemented"; return "" }

func (b *ByteBuffer) Reset() { _ = "STUB: not implemented"; return }
