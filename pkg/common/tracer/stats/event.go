package stats

import (
	"sync"

	"github.com/cloudwego/hertz/pkg/common/errors"
)

type EventIndex int

type Level int

const (
	LevelDisabled Level = iota
	LevelBase
	LevelDetailed
)

type Event interface {
	Index() EventIndex
	Level() Level
}

type event struct {
	idx   EventIndex
	level Level
}

func (e event) Index() EventIndex { _ = "STUB: not implemented"; return *new(EventIndex) }

func (e event) Level() Level { _ = "STUB: not implemented"; return *new(Level) }

const (
	_ EventIndex = iota
	serverHandleStart
	serverHandleFinish
	httpStart
	httpFinish
	readHeaderStart
	readHeaderFinish
	readBodyStart
	readBodyFinish
	writeStart
	writeFinish
	predefinedEventNum
)

var (
	HTTPStart  = newEvent(httpStart, LevelBase)
	HTTPFinish = newEvent(httpFinish, LevelBase)

	ServerHandleStart  = newEvent(serverHandleStart, LevelDetailed)
	ServerHandleFinish = newEvent(serverHandleFinish, LevelDetailed)
	ReadHeaderStart    = newEvent(readHeaderStart, LevelDetailed)
	ReadHeaderFinish   = newEvent(readHeaderFinish, LevelDetailed)
	ReadBodyStart      = newEvent(readBodyStart, LevelDetailed)
	ReadBodyFinish     = newEvent(readBodyFinish, LevelDetailed)
	WriteStart         = newEvent(writeStart, LevelDetailed)
	WriteFinish        = newEvent(writeFinish, LevelDetailed)
)

var (
	ErrNotAllowed = errors.NewPublic("event definition is not allowed after initialization")
	ErrDuplicated = errors.NewPublic("event name is already defined")
)

var (
	lock        sync.RWMutex
	inited      int32
	userDefined = make(map[string]Event)
	maxEventNum = int(predefinedEventNum)
)

func FinishInitialization() { _ = "STUB: not implemented"; return }

func DefineNewEvent(name string, level Level) (Event, error) {
	_ = "STUB: not implemented"
	return *new(Event), nil
}

func MaxEventNum() int { _ = "STUB: not implemented"; return 0 }

func newEvent(idx EventIndex, level Level) Event { _ = "STUB: not implemented"; return *new(Event) }
