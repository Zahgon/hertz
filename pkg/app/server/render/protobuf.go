package render

import (
	"github.com/cloudwego/hertz/pkg/protocol"
)

type ProtoBuf struct {
	Data interface{}
}

var protobufContentType = "application/x-protobuf"

func (r ProtoBuf) Render(resp *protocol.Response) error { _ = "STUB: not implemented"; return nil }

func (r ProtoBuf) WriteContentType(resp *protocol.Response) { _ = "STUB: not implemented"; return }
