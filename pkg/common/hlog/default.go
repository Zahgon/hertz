package hlog

import (
	"context"
	"io"
	"log"
)

func Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

func Notice(v ...interface{}) { _ = "STUB: not implemented"; return }

func Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func Trace(v ...interface{}) { _ = "STUB: not implemented"; return }

func Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Noticef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Tracef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func CtxInfof(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func CtxTracef(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

type defaultLogger struct {
	stdlog *log.Logger
	level  Level
	depth  int
}

func (ll *defaultLogger) SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) SetLevel(lv Level) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) logf(lv Level, format *string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) Fatal(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Error(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Warn(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Notice(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Info(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Debug(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Trace(v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Fatalf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Noticef(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) Tracef(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func (ll *defaultLogger) CtxFatalf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) CtxErrorf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) CtxWarnf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) CtxNoticef(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) CtxInfof(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) CtxDebugf(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (ll *defaultLogger) CtxTracef(ctx context.Context, format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}
