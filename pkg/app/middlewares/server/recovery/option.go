package recovery

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type (
	options struct {
		recoveryHandler func(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte)
	}

	Option func(o *options)
)

func defaultRecoveryHandler(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte) {
	_ = "STUB: not implemented"
	return
}

func newOptions(opts ...Option) *options { _ = "STUB: not implemented"; return nil }

func WithRecoveryHandler(f func(c context.Context, ctx *app.RequestContext, err interface{}, stack []byte)) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}
