package render

import (
	"github.com/cloudwego/hertz/pkg/protocol"
)

type XML struct {
	Data interface{}
}

var xmlContentType = "application/xml; charset=utf-8"

func (r XML) Render(resp *protocol.Response) error { _ = "STUB: not implemented"; return nil }

func (r XML) WriteContentType(w *protocol.Response) { _ = "STUB: not implemented"; return }
