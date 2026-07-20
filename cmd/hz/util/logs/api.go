package logs

func init() {
	defaultLogger = NewStdLogger(LevelInfo)
}

func SetLogger(logger Logger) { _ = "STUB: not implemented"; return }

const (
	LevelDebug = 1 + iota
	LevelInfo
	LevelWarn
	LevelError
)

type Logger interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Warnf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Flush()
	SetLevel(level int) error
}

var defaultLogger Logger

func Errorf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warnf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Infof(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debugf(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Error(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Warn(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Info(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Debug(format string, v ...interface{}) { _ = "STUB: not implemented"; return }

func Flush() { _ = "STUB: not implemented"; return }

func SetLevel(level int) { _ = "STUB: not implemented"; return }
