package stackless

import (
	"sync"
)

func NewFunc(f func(ctx interface{})) func(ctx interface{}) bool {
	_ = "STUB: not implemented"
	return nil
}

func funcWorker(funcWorkCh <-chan *funcWork, f func(ctx interface{})) {
	_ = "STUB: not implemented"
	return
}

func getFuncWork() *funcWork { _ = "STUB: not implemented"; return nil }

func putFuncWork(fw *funcWork) { _ = "STUB: not implemented"; return }

var funcWorkPool sync.Pool

type funcWork struct {
	ctx  interface{}
	done chan struct{}
}
