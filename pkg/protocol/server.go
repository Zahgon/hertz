package protocol

import (
	"context"

	"github.com/cloudwego/hertz/pkg/network"
)

type Server interface {
	Serve(c context.Context, conn network.Conn) error
}

type StreamServer interface {
	Serve(c context.Context, conn network.StreamConn) error
}
