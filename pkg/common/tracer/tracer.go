package tracer

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
)

type Tracer interface {
	Start(ctx context.Context, c *app.RequestContext) context.Context
	Finish(ctx context.Context, c *app.RequestContext)
}

type Controller interface {
	Append(col Tracer)
	DoStart(ctx context.Context, c *app.RequestContext) context.Context
	DoFinish(ctx context.Context, c *app.RequestContext, err error)
	HasTracer() bool
}
