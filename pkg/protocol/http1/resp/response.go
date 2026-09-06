package resp

import (
	"fmt"
	"io"
	"sync"

	"github.com/cloudwego/hertz/pkg/network"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type ErrBodyStreamWritePanic struct {
	error
}

type h1Response struct {
	*protocol.Response
}

func (h1Resp *h1Response) String() string { _ = "STUB: not implemented"; return "" }

func GetHTTP1Response(resp *protocol.Response) fmt.Stringer {
	_ = "STUB: not implemented"
	return *new(fmt.Stringer)
}

func ReadHeaders(resp *protocol.Response, r network.Reader) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadHeaderAndLimitBody(resp *protocol.Response, r network.Reader, maxBodySize int) error {
	_ = "STUB: not implemented"
	return nil
}

func ReadRespBody(resp *protocol.Response, r network.Reader, maxBodySize int) (err error) {
	_ = "STUB: not implemented"
	return nil
}

type clientRespStream struct {
	mu sync.Mutex

	r             io.Reader
	closeCallback func(shouldClose bool) error
}

func (c *clientRespStream) ForceClose() (err error) { _ = "STUB: not implemented"; return nil }

func (c *clientRespStream) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (c *clientRespStream) Read(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

var clientRespStreamPool = sync.Pool{
	New: func() interface{} {
		return &clientRespStream{}
	},
}

func convertClientRespStream(bs io.Reader, fn func(shouldClose bool) error) *clientRespStream {
	_ = "STUB: not implemented"
	return nil
}

func ReadHeaderBodyStream(resp *protocol.Response, r network.Reader,
	maxBodySize int, closeCallBack func(shouldClose bool) error) error {
	_ = "STUB: not implemented"
	return nil
}

var ReadBodyStream = ReadHeaderBodyStream

func ReadRespBodyStream(resp *protocol.Response, r network.Reader,
	maxBodySize int, closeCallBack func(shouldClose bool) error) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func Read(resp *protocol.Response, r network.Reader) error { _ = "STUB: not implemented"; return nil }

func Write(resp *protocol.Response, w network.Writer) error { _ = "STUB: not implemented"; return nil }

func writeBodyStream(resp *protocol.Response, w network.Writer, sendBody bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}
