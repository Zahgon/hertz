package factory

import (
	"github.com/cloudwego/hertz/pkg/protocol/client"
	"github.com/cloudwego/hertz/pkg/protocol/http1"
	"github.com/cloudwego/hertz/pkg/protocol/suite"
)

var _ suite.ClientFactory = (*clientFactory)(nil)

type clientFactory struct {
	option *http1.ClientOptions
}

func (s *clientFactory) NewHostClient() (client client.HostClient, err error) {
	_ = "STUB: not implemented"
	return *new(client.HostClient), nil
}

func NewClientFactory(option *http1.ClientOptions) suite.ClientFactory {
	_ = "STUB: not implemented"
	return *new(suite.ClientFactory)
}
