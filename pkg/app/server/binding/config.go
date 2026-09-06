package binding

import (
	"reflect"

	inDecoder "github.com/cloudwego/hertz/pkg/app/server/binding/internal/decoder"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/cloudwego/hertz/pkg/route/param"
)

type BindConfig struct {
	LooseZeroMode bool

	DisableDefaultTag bool

	DisableStructFieldResolve bool

	EnableDecoderUseNumber bool

	EnableDecoderDisallowUnknownFields bool

	TypeUnmarshalFuncs map[reflect.Type]inDecoder.CustomizeDecodeFunc

	Validator StructValidator

	ValidatorFunc func(req *protocol.Request, v any) error
}

func NewBindConfig() *BindConfig { _ = "STUB: not implemented"; return nil }

func (config *BindConfig) RegTypeUnmarshal(t reflect.Type, fn inDecoder.CustomizeDecodeFunc) error {
	_ = "STUB: not implemented"
	return nil
}

func (config *BindConfig) MustRegTypeUnmarshal(t reflect.Type, fn func(req *protocol.Request, params param.Params, text string) (reflect.Value, error)) {
	_ = "STUB: not implemented"
	return
}

func (config *BindConfig) initTypeUnmarshal() { _ = "STUB: not implemented"; return }

func (config *BindConfig) UseThirdPartyJSONUnmarshaler(fn func(data []byte, v interface{}) error) {
	_ = "STUB: not implemented"
	return
}

func (config *BindConfig) UseStdJSONUnmarshaler() { _ = "STUB: not implemented"; return }

type ValidateErrFactory func(fieldSelector, msg string) error

type ValidateConfig struct {
	ValidateTag string
	ErrFactory  ValidateErrFactory
}

func NewValidateConfig() *ValidateConfig { _ = "STUB: not implemented"; return nil }

func (config *ValidateConfig) MustRegValidateFunc(funcName string, fn func(args ...interface{}) error, force ...bool) {
	_ = "STUB: not implemented"
	return
}

func (config *ValidateConfig) SetValidatorErrorFactory(errFactory ValidateErrFactory) {
	_ = "STUB: not implemented"
	return
}

func (config *ValidateConfig) SetValidatorTag(tag string) { _ = "STUB: not implemented"; return }
