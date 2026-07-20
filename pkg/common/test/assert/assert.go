package assert

import (
	"reflect"
)

type testingT interface {
	Helper()
	Fatal(args ...any)
	Fatalf(format string, args ...any)
}

func Assert(t testingT, cond bool, val ...interface{}) { _ = "STUB: not implemented"; return }

func Assertf(t testingT, cond bool, format string, val ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func DeepEqual(t testingT, expected, actual interface{}) { _ = "STUB: not implemented"; return }

func isNil(rv reflect.Value) bool { _ = "STUB: not implemented"; return false }

func Nil(t testingT, data interface{}) { _ = "STUB: not implemented"; return }

func NotNil(t testingT, data interface{}) { _ = "STUB: not implemented"; return }

func NotEqual(t testingT, expected, actual interface{}) { _ = "STUB: not implemented"; return }

func True(t testingT, obj interface{}) { _ = "STUB: not implemented"; return }

func False(t testingT, obj interface{}) { _ = "STUB: not implemented"; return }

func Panic(t testingT, fn func()) { _ = "STUB: not implemented"; return }

func NotPanic(t testingT, fn func()) { _ = "STUB: not implemented"; return }
