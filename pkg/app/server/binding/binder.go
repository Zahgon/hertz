package binding

import (
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/route/param"
)

type Binder interface {
	Name() string
	Bind(*protocol.Request, interface{}, param.Params) error
	BindQuery(*protocol.Request, interface{}) error
	BindHeader(*protocol.Request, interface{}) error
	BindPath(*protocol.Request, interface{}, param.Params) error
	BindForm(*protocol.Request, interface{}) error
	BindJSON(*protocol.Request, interface{}) error
	BindProtobuf(*protocol.Request, interface{}) error

	Validate(*protocol.Request, interface{}) error
}
