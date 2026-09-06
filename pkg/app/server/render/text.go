package render

import (
	"github.com/cloudwego/hertz/pkg/protocol"
)

type String struct {
	Format string
	Data   []interface{}
}

var plainContentType = "text/plain; charset=utf-8"

func (r String) Render(resp *protocol.Response) error { _ = "STUB: not implemented"; return nil }

func (r String) WriteContentType(resp *protocol.Response) { _ = "STUB: not implemented"; return }
