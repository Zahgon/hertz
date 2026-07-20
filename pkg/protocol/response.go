package protocol

import (
	"io"
	"net"
	"sync"

	"github.com/cloudwego/hertz/internal/nocopy"
	"github.com/cloudwego/hertz/pkg/common/bytebufferpool"
	"github.com/cloudwego/hertz/pkg/network"
)

var (
	responsePool sync.Pool

	NoResponseBody = noBody{}
)

type Response struct {
	noCopy nocopy.NoCopy //lint:ignore U1000 until noCopy is used

	Header ResponseHeader

	ImmediateHeaderFlush bool

	bodyStream      io.Reader
	w               responseBodyWriter
	body            *bytebufferpool.ByteBuffer
	bodyRaw         []byte
	maxKeepBodySize int

	SkipBody bool

	raddr net.Addr

	laddr net.Addr

	hijackWriter network.ExtWriter
}

func (resp *Response) GetHijackWriter() network.ExtWriter {
	_ = "STUB: not implemented"
	return *new(network.ExtWriter)
}

func (resp *Response) HijackWriter(writer network.ExtWriter) { _ = "STUB: not implemented"; return }

type responseBodyWriter struct {
	r *Response
}

func (w *responseBodyWriter) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (resp *Response) MustSkipBody() bool { _ = "STUB: not implemented"; return false }

func (resp *Response) BodyGunzip() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (resp *Response) SetConnectionClose() { _ = "STUB: not implemented"; return }

func (resp *Response) SetBodyString(body string) { _ = "STUB: not implemented"; return }

//nolint:errcheck
//nolint:errcheck

func (resp *Response) ConstructBodyStream(body *bytebufferpool.ByteBuffer, bodyStream io.Reader) {
	_ = "STUB: not implemented"
	return
}

func (resp *Response) BodyWriter() io.Writer { _ = "STUB: not implemented"; return *new(io.Writer) }

func (resp *Response) SetStatusCode(statusCode int) { _ = "STUB: not implemented"; return }

func (resp *Response) SetMaxKeepBodySize(n int) { _ = "STUB: not implemented"; return }

func (resp *Response) BodyBytes() []byte { _ = "STUB: not implemented"; return nil }

func (resp *Response) HasBodyBytes() bool { _ = "STUB: not implemented"; return false }

func (resp *Response) CopyToSkipBody(dst *Response) { _ = "STUB: not implemented"; return }

func (resp *Response) IsBodyStream() bool { _ = "STUB: not implemented"; return false }

func (resp *Response) SetBodyStream(bodyStream io.Reader, bodySize int) {
	_ = "STUB: not implemented"
	return
}

func (resp *Response) SetBodyStreamNoReset(bodyStream io.Reader, bodySize int) {
	_ = "STUB: not implemented"
	return
}

func (resp *Response) BodyE() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

//nolint:errcheck

func (resp *Response) Body() []byte { _ = "STUB: not implemented"; return nil }

func (resp *Response) BodyWriteTo(w io.Writer) error { _ = "STUB: not implemented"; return nil }

//nolint:errcheck

//nolint:errcheck

func (resp *Response) CopyTo(dst *Response) { _ = "STUB: not implemented"; return }

func SwapResponseBody(a, b *Response) { _ = "STUB: not implemented"; return }

func (resp *Response) Reset() { _ = "STUB: not implemented"; return }

func (resp *Response) resetSkipHeader() { _ = "STUB: not implemented"; return }

func (resp *Response) ResetBody() { _ = "STUB: not implemented"; return }

//nolint:errcheck

func (resp *Response) SetBodyRaw(body []byte) { _ = "STUB: not implemented"; return }

func (resp *Response) StatusCode() int { _ = "STUB: not implemented"; return 0 }

func (resp *Response) SetBody(body []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func (resp *Response) BodyStream() io.Reader { _ = "STUB: not implemented"; return *new(io.Reader) }

func (resp *Response) Hijack() (network.Conn, error) {
	_ = "STUB: not implemented"
	return *new(network.Conn), nil
}

func (resp *Response) AppendBody(p []byte) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func (resp *Response) AppendBodyString(s string) { _ = "STUB: not implemented"; return }

//nolint:errcheck

//nolint:errcheck

//nolint:errcheck

func (resp *Response) ConnectionClose() bool { _ = "STUB: not implemented"; return false }

func (resp *Response) CloseBodyStream() error { _ = "STUB: not implemented"; return nil }

func (resp *Response) BodyBuffer() *bytebufferpool.ByteBuffer {
	_ = "STUB: not implemented"
	return nil
}

func gunzipData(p []byte) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (resp *Response) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (resp *Response) LocalAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (resp *Response) ParseNetAddr(conn network.Conn) { _ = "STUB: not implemented"; return }

func AcquireResponse() *Response { _ = "STUB: not implemented"; return nil }

func ReleaseResponse(resp *Response) { _ = "STUB: not implemented"; return }
