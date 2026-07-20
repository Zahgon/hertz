package sd

import (
	"github.com/cloudwego/hertz/pkg/app/client"
	"github.com/cloudwego/hertz/pkg/app/client/discovery"
)

func Discovery(resolver discovery.Resolver, opts ...ServiceDiscoveryOption) client.Middleware {
	_ = "STUB: not implemented"
	return *new(client.Middleware)
}
