package adaptor

import (
	"net/http"

	"github.com/cloudwego/hertz/pkg/protocol"
)

type compatResponse struct {
	h           *protocol.Response
	header      http.Header
	writeHeader bool
}

func (c *compatResponse) Header() http.Header { _ = "STUB: not implemented"; return *new(http.Header) }

func (c *compatResponse) Write(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *compatResponse) WriteHeader(statusCode int) { _ = "STUB: not implemented"; return }

func GetCompatResponseWriter(resp *protocol.Response) http.ResponseWriter {
	_ = "STUB: not implemented"
	return *new(http.ResponseWriter)
}
