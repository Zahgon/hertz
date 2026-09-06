package generator

import (
	"text/template"
)

var DefaultDelimiters = [2]string{"{{", "}}"}

type TemplateConfig struct {
	Layouts []Template `yaml:"layouts"`
}

const (
	Skip   = "skip"
	Cover  = "cover"
	Append = "append"
)

type Template struct {
	Default        bool
	Path           string         `yaml:"path"`
	Delims         [2]string      `yaml:"delims"`
	Body           string         `yaml:"body"`
	Disable        bool           `yaml:"disable"`
	LoopMethod     bool           `yaml:"loop_method"`
	LoopService    bool           `yaml:"loop_service"`
	UpdateBehavior UpdateBehavior `yaml:"update_behavior"`
}

type UpdateBehavior struct {
	Type string `yaml:"type"`

	AppendKey      string   `yaml:"append_key"`
	InsertKey      string   `yaml:"insert_key"`
	AppendTpl      string   `yaml:"append_content_tpl"`
	ImportTpl      []string `yaml:"import_tpl"`
	AppendLocation string   `yaml:"append_location"`
}

type TemplateGenerator struct {
	OutputDir    string
	Config       *TemplateConfig
	Excludes     []string
	tpls         map[string]*template.Template
	tplsInfo     map[string]*Template
	dirs         map[string]bool
	isPackageTpl bool

	files         []File
	excludedFiles map[string]*File
}

func (tg *TemplateGenerator) Init() error { _ = "STUB: not implemented"; return nil }

func (tg *TemplateGenerator) loadLayout(layout Template, tplName string, isDefaultTpl bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (tg *TemplateGenerator) Generate(input interface{}, tplName, filepath string, noRepeat bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (tg *TemplateGenerator) Persist() error { _ = "STUB: not implemented"; return nil }

func (tg *TemplateGenerator) GetFormatAndExcludedFiles() ([]File, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (tg *TemplateGenerator) Files() []File { _ = "STUB: not implemented"; return nil }

func (tg *TemplateGenerator) Degenerate() error { _ = "STUB: not implemented"; return nil }
