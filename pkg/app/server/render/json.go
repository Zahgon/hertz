package render

import (
	hjson "github.com/cloudwego/hertz/pkg/common/json"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type JSONMarshaler func(v interface{}) ([]byte, error)

var jsonMarshalFunc JSONMarshaler

func init() {
	ResetJSONMarshal(hjson.Marshal)
}

func ResetJSONMarshal(fn JSONMarshaler) { _ = "STUB: not implemented"; return }

func ResetStdJSONMarshal() { _ = "STUB: not implemented"; return }

type JSONRender struct {
	Data interface{}
}

var jsonContentType = "application/json; charset=utf-8"

func (r JSONRender) Render(resp *protocol.Response) error { _ = "STUB: not implemented"; return nil }

func (r JSONRender) WriteContentType(resp *protocol.Response) { _ = "STUB: not implemented"; return }

type PureJSON struct {
	Data interface{}
}

func (r PureJSON) Render(resp *protocol.Response) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r PureJSON) WriteContentType(resp *protocol.Response) { _ = "STUB: not implemented"; return }

type IndentedJSON struct {
	Data interface{}
}

func (r IndentedJSON) Render(resp *protocol.Response) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (r IndentedJSON) WriteContentType(resp *protocol.Response) { _ = "STUB: not implemented"; return }
