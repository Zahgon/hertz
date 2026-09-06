package server

import (
	"github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/route"
)

type Hertz struct {
	*route.Engine
	signalWaiter func(err chan error) error
}

func New(opts ...config.Option) *Hertz { _ = "STUB: not implemented"; return nil }

func Default(opts ...config.Option) *Hertz { _ = "STUB: not implemented"; return nil }

func (h *Hertz) Spin() { _ = "STUB: not implemented"; return }

func (h *Hertz) SetCustomSignalWaiter(f func(err chan error) error) {
	_ = "STUB: not implemented"
	return
}

func waitSignal(errCh chan error) error { _ = "STUB: not implemented"; return nil }

func (h *Hertz) initOnRunHooks(errChan chan error) { _ = "STUB: not implemented"; return }
