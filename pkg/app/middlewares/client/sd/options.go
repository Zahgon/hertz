package sd

import (
	"github.com/cloudwego/hertz/pkg/app/client/discovery"
	"github.com/cloudwego/hertz/pkg/app/client/loadbalance"
)

type ServiceDiscoveryOptions struct {
	Resolver discovery.Resolver

	Balancer loadbalance.Loadbalancer

	LbOpts loadbalance.Options
}

func (o *ServiceDiscoveryOptions) Apply(opts []ServiceDiscoveryOption) {
	_ = "STUB: not implemented"
	return
}

type ServiceDiscoveryOption struct {
	F func(o *ServiceDiscoveryOptions)
}

func WithCustomizedAddrs(addrs ...string) ServiceDiscoveryOption {
	_ = "STUB: not implemented"
	return *new(ServiceDiscoveryOption)
}

func WithLoadBalanceOptions(lb loadbalance.Loadbalancer, options loadbalance.Options) ServiceDiscoveryOption {
	_ = "STUB: not implemented"
	return *new(ServiceDiscoveryOption)
}
