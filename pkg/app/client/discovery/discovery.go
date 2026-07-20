package discovery

import (
	"context"
	"net"
)

type TargetInfo struct {
	Host string
	Tags map[string]string
}

type Resolver interface {
	Target(ctx context.Context, target *TargetInfo) string

	Resolve(ctx context.Context, desc string) (Result, error)

	Name() string
}

type SynthesizedResolver struct {
	TargetFunc  func(ctx context.Context, target *TargetInfo) string
	ResolveFunc func(ctx context.Context, key string) (Result, error)
	NameFunc    func() string
}

func (sr SynthesizedResolver) Target(ctx context.Context, target *TargetInfo) string {
	_ = "STUB: not implemented"
	return ""
}

func (sr SynthesizedResolver) Resolve(ctx context.Context, key string) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

func (sr SynthesizedResolver) Name() string { _ = "STUB: not implemented"; return "" }

type Instance interface {
	Address() net.Addr
	Weight() int
	Tag(key string) (value string, exist bool)
}

type instance struct {
	addr   net.Addr
	weight int
	tags   map[string]string
}

func (i *instance) Address() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (i *instance) Weight() int { _ = "STUB: not implemented"; return 0 }

func (i *instance) Tag(key string) (value string, exist bool) {
	_ = "STUB: not implemented"
	return "", false
}

func NewInstance(network, address string, weight int, tags map[string]string) Instance {
	_ = "STUB: not implemented"
	return *new(Instance)
}

type Result struct {
	CacheKey  string
	Instances []Instance
}
