package logs

import (
	"bytes"
	"log"
)

type StdLogger struct {
	level      int
	outLogger  *log.Logger
	warnLogger *log.Logger
	errLogger  *log.Logger
	out        *bytes.Buffer
	warn       *bytes.Buffer
	err        *bytes.Buffer
	Defer      bool
	ErrOnly    bool
}

func NewStdLogger(level int) *StdLogger { _ = "STUB: not implemented"; return nil }

func (stdLogger *StdLogger) Debugf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (stdLogger *StdLogger) Infof(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (stdLogger *StdLogger) Warnf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (stdLogger *StdLogger) Errorf(format string, v ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func (stdLogger *StdLogger) Flush() { _ = "STUB: not implemented"; return }

func (stdLogger *StdLogger) FlushOut() { _ = "STUB: not implemented"; return }

func (stdLogger *StdLogger) Err() string { _ = "STUB: not implemented"; return "" }

func (stdLogger *StdLogger) Warn() string { _ = "STUB: not implemented"; return "" }

func (stdLogger *StdLogger) FlushErr() { _ = "STUB: not implemented"; return }

func (stdLogger *StdLogger) OutLines() []string { _ = "STUB: not implemented"; return nil }

func (stdLogger *StdLogger) Out() []byte { _ = "STUB: not implemented"; return nil }

func (stdLogger *StdLogger) SetLevel(level int) error { _ = "STUB: not implemented"; return nil }
