package client

import (
	"context"

	"github.com/cloudwego/hertz/pkg/protocol"
)

type Endpoint func(ctx context.Context, req *protocol.Request, resp *protocol.Response) (err error)

type Middleware func(Endpoint) Endpoint

func chain(mws ...Middleware) Middleware { _ = "STUB: not implemented"; return *new(Middleware) }
