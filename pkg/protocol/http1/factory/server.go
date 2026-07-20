package factory

import (
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/protocol/http1"
	"github.com/cloudwego/hertz/pkg/protocol/suite"
)

var _ suite.ServerFactory = (*serverFactory)(nil)

type serverFactory struct {
	option *http1.Option
}

func (s *serverFactory) New(core suite.Core) (server protocol.Server, err error) {
	_ = "STUB: not implemented"
	return *new(protocol.Server), nil
}

func NewServerFactory(option *http1.Option) suite.ServerFactory {
	_ = "STUB: not implemented"
	return *new(suite.ServerFactory)
}
