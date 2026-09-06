package loadbalance

import (
	"sync"

	"github.com/cloudwego/hertz/pkg/app/client/discovery"
	"golang.org/x/sync/singleflight"
)

type weightedBalancer struct {
	cachedWeightInfo sync.Map
	sfg              singleflight.Group
}

type weightInfo struct {
	instances []discovery.Instance
	entries   []int
	weightSum int
}

func NewWeightedBalancer() Loadbalancer { _ = "STUB: not implemented"; return *new(Loadbalancer) }

func (wb *weightedBalancer) calcWeightInfo(e discovery.Result) *weightInfo {
	_ = "STUB: not implemented"
	return nil
}

func (wb *weightedBalancer) Pick(e discovery.Result) discovery.Instance {
	_ = "STUB: not implemented"
	return *new(discovery.Instance)
}

func (wb *weightedBalancer) Rebalance(e discovery.Result) { _ = "STUB: not implemented"; return }

func (wb *weightedBalancer) Delete(cacheKey string) { _ = "STUB: not implemented"; return }

func (wb *weightedBalancer) Name() string { _ = "STUB: not implemented"; return "" }
