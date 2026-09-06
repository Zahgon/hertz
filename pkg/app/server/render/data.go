package render

import "github.com/cloudwego/hertz/pkg/protocol"

type Data struct {
	ContentType string
	Data        []byte
}

func (r Data) Render(resp *protocol.Response) (err error) { _ = "STUB: not implemented"; return nil }

func (r Data) WriteContentType(resp *protocol.Response) { _ = "STUB: not implemented"; return }
