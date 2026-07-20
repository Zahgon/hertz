package traceinfo

import (
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/common/tracer/stats"
)

var _ HTTPStats = (*httpStats)(nil)

var (
	eventPool   sync.Pool
	once        sync.Once
	maxEventNum int
)

type event struct {
	event  stats.Event
	status stats.Status
	info   string
	time   time.Time
}

func (e *event) Event() stats.Event { _ = "STUB: not implemented"; return *new(stats.Event) }

func (e *event) Status() stats.Status { _ = "STUB: not implemented"; return *new(stats.Status) }

func (e *event) Info() string { _ = "STUB: not implemented"; return "" }

func (e *event) Time() time.Time { _ = "STUB: not implemented"; return *new(time.Time) }

func (e *event) IsNil() bool { _ = "STUB: not implemented"; return false }

func newEvent() interface{} { _ = "STUB: not implemented"; return nil }

func (e *event) zero() { _ = "STUB: not implemented"; return }

func (e *event) Recycle() { _ = "STUB: not implemented"; return }

type httpStats struct {
	sync.RWMutex
	level stats.Level

	eventMap []Event

	sendSize int
	recvSize int

	err      error
	panicErr interface{}
}

func init() {
	eventPool.New = newEvent
}

func (h *httpStats) Record(e stats.Event, status stats.Status, info string) {
	_ = "STUB: not implemented"
	return
}

func (h *httpStats) SendSize() int { _ = "STUB: not implemented"; return 0 }

func (h *httpStats) RecvSize() int { _ = "STUB: not implemented"; return 0 }

func (h *httpStats) Error() error { _ = "STUB: not implemented"; return nil }

func (h *httpStats) Panicked() (bool, interface{}) { _ = "STUB: not implemented"; return false, nil }

func (h *httpStats) GetEvent(e stats.Event) Event { _ = "STUB: not implemented"; return *new(Event) }

func (h *httpStats) Level() stats.Level { _ = "STUB: not implemented"; return *new(stats.Level) }

func (h *httpStats) SetSendSize(size int) { _ = "STUB: not implemented"; return }

func (h *httpStats) SetRecvSize(size int) { _ = "STUB: not implemented"; return }

func (h *httpStats) SetError(err error) { _ = "STUB: not implemented"; return }

func (h *httpStats) SetPanicked(x interface{}) { _ = "STUB: not implemented"; return }

func (h *httpStats) SetLevel(level stats.Level) { _ = "STUB: not implemented"; return }

func (h *httpStats) Reset() { _ = "STUB: not implemented"; return }

func (h *httpStats) ImmutableView() HTTPStats { _ = "STUB: not implemented"; return *new(HTTPStats) }

func NewHTTPStats() HTTPStats { _ = "STUB: not implemented"; return *new(HTTPStats) }
