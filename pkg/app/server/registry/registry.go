package registry

import "net"

const (
	DefaultWeight = 10
)

type Registry interface {
	Register(info *Info) error
	Deregister(info *Info) error
}

type Info struct {
	ServiceName string

	Addr net.Addr

	Weight int

	Tags map[string]string
}

var NoopRegistry Registry = &noopRegistry{}

type noopRegistry struct{}

func (e noopRegistry) Register(*Info) error { _ = "STUB: not implemented"; return nil }

func (e noopRegistry) Deregister(*Info) error { _ = "STUB: not implemented"; return nil }
