package recovery

import (
	"github.com/cloudwego/hertz/pkg/app"
)

var (
	dunno     = []byte("???")
	centerDot = []byte("·")
	dot       = []byte(".")
	slash     = []byte("/")
)

func Recovery(opts ...Option) app.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(app.HandlerFunc)
}

func stack(skip int) []byte { _ = "STUB: not implemented"; return nil }

func source(lines [][]byte, n int) []byte { _ = "STUB: not implemented"; return nil }

func function(pc uintptr) []byte { _ = "STUB: not implemented"; return nil }
