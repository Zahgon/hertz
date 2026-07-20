package hlog

import (
	"io"
	"log"
	"os"
)

var (
	logger FullLogger = &defaultLogger{
		stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
		depth:  4,
	}

	sysLogger FullLogger = &systemLogger{
		&defaultLogger{
			stdlog: log.New(os.Stderr, "", log.LstdFlags|log.Lshortfile|log.Lmicroseconds),
			depth:  4,
		},
		systemLogPrefix,
	}
)

func SetOutput(w io.Writer) { _ = "STUB: not implemented"; return }

func SetLevel(lv Level) { _ = "STUB: not implemented"; return }

func DefaultLogger() FullLogger { _ = "STUB: not implemented"; return *new(FullLogger) }

func SystemLogger() FullLogger { _ = "STUB: not implemented"; return *new(FullLogger) }

func SetSystemLogger(v FullLogger) { _ = "STUB: not implemented"; return }

func SetLogger(v FullLogger) { _ = "STUB: not implemented"; return }
