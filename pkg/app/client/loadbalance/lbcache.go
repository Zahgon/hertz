package loadbalance

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/cloudwego/hertz/pkg/app/client/discovery"
	"github.com/cloudwego/hertz/pkg/protocol"
	"golang.org/x/sync/singleflight"
)

type cacheResult struct {
	res         atomic.Value
	expire      int32
	serviceName string
}

var (
	balancerFactories    sync.Map
	balancerFactoriesSfg singleflight.Group
)

func cacheKey(resolver, balancer string, opts Options) string { _ = "STUB: not implemented"; return "" }

type BalancerFactory struct {
	opts     Options
	cache    sync.Map
	resolver discovery.Resolver
	impl     Loadbalancer
	implCtx  LoadbalancerCtx
	sfg      singleflight.Group
}

type Config struct {
	Resolver discovery.Resolver
	Balancer Loadbalancer
	LbOpts   Options
}

func NewBalancerFactory(config Config) *BalancerFactory { _ = "STUB: not implemented"; return nil }

func (b *BalancerFactory) watcher() { _ = "STUB: not implemented"; return }

func renameResultCacheKey(res *discovery.Result, resolverName string) {
	_ = "STUB: not implemented"
	return
}

func (b *BalancerFactory) refresh() { _ = "STUB: not implemented"; return }

func (b *BalancerFactory) GetInstance(ctx context.Context, req *protocol.Request) (discovery.Instance, error) {
	_ = "STUB: not implemented"
	return *new(discovery.Instance), nil
}

func (b *BalancerFactory) getCacheResult(ctx context.Context, req *protocol.Request) (*cacheResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
