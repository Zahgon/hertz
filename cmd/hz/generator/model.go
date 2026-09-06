package generator

import (
	"text/template"

	"github.com/cloudwego/hertz/cmd/hz/generator/model"
	"github.com/cloudwego/hertz/cmd/hz/meta"
)

type Option string

const (
	OptionMarshalEnumToText  Option = "MarshalEnumToText"
	OptionTypedefAsTypeAlias Option = "TypedefAsTypeAlias"
)

type Backend interface {
	Template() (*template.Template, error)
	List() map[string]string
	SetOption(opts string) error
	GetOptions() []string
	Funcs(name string, fn interface{}) error
}

type GolangBackend struct{}

func (gb *GolangBackend) Template() (*template.Template, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (gb *GolangBackend) List() map[string]string { _ = "STUB: not implemented"; return nil }

func (gb *GolangBackend) SetOption(opts string) error { _ = "STUB: not implemented"; return nil }

func (gb *GolangBackend) GetOptions() []string { _ = "STUB: not implemented"; return nil }

func (gb *GolangBackend) Funcs(name string, fn interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func switchBackend(backend meta.Backend) Backend { _ = "STUB: not implemented"; return *new(Backend) }

func loadThirdPartyBackend(plugin string) Backend { _ = "STUB: not implemented"; return *new(Backend) }

func (pkgGen *HttpPackageGenerator) LoadBackend(backend meta.Backend) error {
	_ = "STUB: not implemented"
	return nil
}

func (pkgGen *HttpPackageGenerator) GenModel(data *model.Model, gen bool) error {
	_ = "STUB: not implemented"
	return nil
}

func removeDuplicateImport(data *model.Model) { _ = "STUB: not implemented"; return }
