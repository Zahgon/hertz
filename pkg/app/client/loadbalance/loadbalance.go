package loadbalance

import (
	"context"
	"time"

	"github.com/cloudwego/hertz/pkg/app/client/discovery"
	"github.com/cloudwego/hertz/pkg/protocol"
)

type Loadbalancer interface {
	Pick(discovery.Result) discovery.Instance

	Rebalance(discovery.Result)

	Delete(string)

	Name() string
}

type LoadbalancerCtx interface {
	Loadbalancer

	PickCtx(ctx context.Context, req *protocol.Request, e discovery.Result) discovery.Instance
}

const (
	DefaultRefreshInterval = 5 * time.Second
	DefaultExpireInterval  = 15 * time.Second
)

var DefaultLbOpts = Options{
	RefreshInterval: DefaultRefreshInterval,
	ExpireInterval:  DefaultExpireInterval,
}

type Options struct {
	RefreshInterval time.Duration

	ExpireInterval time.Duration
}

func (v *Options) Check() { _ = "STUB: not implemented"; return }
