package sse

import (
	"sync"
	"time"
)

const (
	fieldID = 1 << iota
	fieldType
	fieldData
	fieldRetry
)

type Event struct {
	ID   string
	Type string
	Data []byte

	Retry time.Duration

	bitset uint8
}

var poolEvent = sync.Pool{}

func NewEvent() *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) Release() { _ = "STUB: not implemented"; return }

func (e *Event) String() string { _ = "STUB: not implemented"; return "" }

func (e *Event) Reset() { _ = "STUB: not implemented"; return }

func (e *Event) Clone() *Event { _ = "STUB: not implemented"; return nil }

func (e *Event) IsSetID() bool { _ = "STUB: not implemented"; return false }

func (e *Event) IsSetType() bool { _ = "STUB: not implemented"; return false }

func (e *Event) IsSetRetry() bool { _ = "STUB: not implemented"; return false }

func (e *Event) IsSetData() bool { _ = "STUB: not implemented"; return false }

func (e *Event) SetID(id string) { _ = "STUB: not implemented"; return }

func (e *Event) SetEvent(eventType string) { _ = "STUB: not implemented"; return }

func (e *Event) SetData(data []byte) { _ = "STUB: not implemented"; return }

func (e *Event) SetDataString(data string) { _ = "STUB: not implemented"; return }

func (e *Event) AppendData(data []byte) { _ = "STUB: not implemented"; return }

func (e *Event) AppendDataString(data string) { _ = "STUB: not implemented"; return }

func (e *Event) SetRetry(retry time.Duration) { _ = "STUB: not implemented"; return }
