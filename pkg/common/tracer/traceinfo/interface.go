package traceinfo

import (
	"time"

	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
)

type HTTPStats interface {
	Record(event stats.Event, status stats.Status, info string)
	GetEvent(event stats.Event) Event
	SendSize() int
	SetSendSize(size int)
	RecvSize() int
	SetRecvSize(size int)
	Error() error
	SetError(err error)
	Panicked() (bool, interface{})
	SetPanicked(x interface{})
	Level() stats.Level
	SetLevel(level stats.Level)
	Reset()
}

type Event interface {
	Event() stats.Event
	Status() stats.Status
	Info() string
	Time() time.Time
	IsNil() bool
}

type TraceInfo interface {
	Stats() HTTPStats
	Reset()
}
