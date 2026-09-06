package ut

import (
	"bytes"

	"github.com/cloudwego/hertz/pkg/protocol"
)

type ResponseRecorder struct {
	Code int

	header *protocol.ResponseHeader

	Body *bytes.Buffer

	Flushed bool

	result      *protocol.Response
	wroteHeader bool
}

func NewRecorder() *ResponseRecorder { _ = "STUB: not implemented"; return nil }

func (rw *ResponseRecorder) Header() *protocol.ResponseHeader {
	_ = "STUB: not implemented"
	return nil
}

func (rw *ResponseRecorder) Write(buf []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *ResponseRecorder) WriteString(str string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rw *ResponseRecorder) WriteHeader(code int) { _ = "STUB: not implemented"; return }

func (rw *ResponseRecorder) Flush() { _ = "STUB: not implemented"; return }

func (rw *ResponseRecorder) Result() *protocol.Response { _ = "STUB: not implemented"; return nil }
