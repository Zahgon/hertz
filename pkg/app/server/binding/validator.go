package binding

import (
	"reflect"
	"sync"

	"github.com/cloudwego/hertz/pkg/protocol"
)

type ValidatorFunc func(*protocol.Request, interface{}) error

type StructValidator interface {
	ValidateStruct(interface{}) error
	Engine() interface{}
	ValidateTag() string
}

var hasValidateTagCache sync.Map

func MakeValidatorFunc(s StructValidator) ValidatorFunc {
	_ = "STUB: not implemented"
	return *new(ValidatorFunc)
}

func containsStructTag(rt reflect.Type, tag string, checking map[reflect.Type]bool) bool {
	_ = "STUB: not implemented"
	return false
}
