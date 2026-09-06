package ut

import (
	"io"

	"github.com/cloudwego/hertz/pkg/route"
)

type Header struct {
	Key   string
	Value string
}

type Body struct {
	Body io.Reader
	Len  int
}

func PerformRequest(engine *route.Engine, method, url string, body *Body, headers ...Header) *ResponseRecorder {
	_ = "STUB: not implemented"
	return nil
}
