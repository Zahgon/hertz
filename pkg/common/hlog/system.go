package hlog

import (
	"context"
	"io"
	"strings"
	"sync"
)

var silentMode = false

func SetSilentMode(s bool) { _ = "STUB: not implemented"; return }

var builderPool = sync.Pool{New: func() interface{} {
	return &strings.Builder{}
}}

type systemLogger struct {
	logger FullLogger
	prefix string
}

func (ll *systemLogger) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) SetLevel(lv Level) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Notice(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Trace(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Noticef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) Tracef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *systemLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *systemLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *systemLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *systemLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *systemLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *systemLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *systemLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *systemLogger) addPrefix(format string) string { _ = "STUB: not implemented"; return "" }
