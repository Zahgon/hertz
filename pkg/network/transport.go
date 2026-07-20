package network

import (
	"context"
)

type Transporter interface {
	Close() error

	Shutdown(ctx context.Context) error

	ListenAndServe(onData OnData) error
}

type OnData func(ctx context.Context, conn interface{}) error
