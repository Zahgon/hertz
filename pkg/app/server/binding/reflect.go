package binding

import (
	"reflect"
	"unsafe"
)

func valueAndTypeID(v interface{}) (reflect.Value, uintptr) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), 0
}

type emptyInterface struct {
	typeID  uintptr
	dataPtr unsafe.Pointer
}

func checkPointer(rv reflect.Value) error { _ = "STUB: not implemented"; return nil }

func dereferenceType(t reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}
