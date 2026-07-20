package render

import "github.com/cloudwego/hertz/pkg/protocol"

type Render interface {
	Render(resp *protocol.Response) error

	WriteContentType(resp *protocol.Response)
}

var (
	_ Render = JSONRender{}
	_ Render = String{}
	_ Render = Data{}
)

func writeContentType(resp *protocol.Response, value string) { _ = "STUB: not implemented"; return }
