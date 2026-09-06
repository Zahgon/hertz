package sse

import (
	"github.com/cloudwego/hertz/pkg/protocol"
)

const LastEventIDHeader = "Last-Event-ID"

func GetLastEventID(req *protocol.Request) string { _ = "STUB: not implemented"; return "" }

func SetLastEventID(req *protocol.Request, id string) { _ = "STUB: not implemented"; return }

func AddAcceptMIME(req *protocol.Request) { _ = "STUB: not implemented"; return }

func sseEventType(v []byte) string { _ = "STUB: not implemented"; return "" }

func hasCRLF(s string) bool { _ = "STUB: not implemented"; return false }

func scanEOL(data []byte, atEOF bool) (advance int, token []byte, err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}
