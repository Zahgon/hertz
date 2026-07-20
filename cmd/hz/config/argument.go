package config

import (
	"github.com/cloudwego/hertz/cmd/hz/meta"
	"github.com/urfave/cli/v2"
)

type Argument struct {
	CmdType        string
	Verbose        bool
	Cwd            string
	OutDir         string
	HandlerDir     string
	ModelDir       string
	RouterDir      string
	ClientDir      string
	BaseDomain     string
	ForceClientDir string

	IdlType       string
	IdlPaths      []string
	RawOptPkg     []string
	OptPkgMap     map[string]string
	Includes      []string
	PkgPrefix     string
	TrimGoPackage string

	Gopath      string
	Gosrc       string
	Gomod       string
	Gopkg       string
	ServiceName string
	Use         string
	NeedGoMod   bool

	JSONEnumStr          bool
	QueryEnumAsInt       bool
	UnsetOmitempty       bool
	ProtobufCamelJSONTag bool
	ProtocOptions        []string
	ThriftOptions        []string
	ProtobufPlugins      []string
	ThriftPlugins        []string
	SnakeName            bool
	RmTags               []string
	Excludes             []string
	NoRecurse            bool
	HandlerByMethod      bool
	ForceNew             bool
	ForceUpdateClient    bool
	SnakeStyleMiddleware bool
	EnableExtends        bool
	SortRouter           bool

	EnableClientOptional bool

	CustomizeLayout     string
	CustomizeLayoutData string
	CustomizePackage    string
	ModelBackend        string
}

func NewArgument() *Argument { _ = "STUB: not implemented"; return nil }

func (arg *Argument) Parse(c *cli.Context, cmd string) (*Argument, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (arg *Argument) parseStringSlice(c *cli.Context) { _ = "STUB: not implemented"; return }

func (arg *Argument) UpdateByManifest(m *meta.Manifest) { _ = "STUB: not implemented"; return }

func (arg *Argument) checkPath() error { _ = "STUB: not implemented"; return nil }

func (arg *Argument) checkIDL() error { _ = "STUB: not implemented"; return nil }

func (arg *Argument) IsUpdate() bool { _ = "STUB: not implemented"; return false }

func (arg *Argument) IsNew() bool { _ = "STUB: not implemented"; return false }

func (arg *Argument) checkPackage() error { _ = "STUB: not implemented"; return nil }

func (arg *Argument) Pack() ([]string, error) { _ = "STUB: not implemented"; return nil, nil }

func (arg *Argument) Unpack(data []string) error { _ = "STUB: not implemented"; return nil }

func (arg *Argument) Fork() *Argument { _ = "STUB: not implemented"; return nil }

func (arg *Argument) GetGoPackage() (string, error) { _ = "STUB: not implemented"; return "", nil }

func IdlTypeToCompiler(idlType string) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (arg *Argument) ModelPackagePrefix() (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (arg *Argument) ModelOutDir() string { _ = "STUB: not implemented"; return "" }

func (arg *Argument) GetHandlerDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (arg *Argument) GetModelDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (arg *Argument) GetRouterDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (arg *Argument) GetClientDir() (string, error) { _ = "STUB: not implemented"; return "", nil }

func (arg *Argument) InitManifest(m *meta.Manifest) { _ = "STUB: not implemented"; return }

func (arg *Argument) UpdateManifest(m *meta.Manifest) { _ = "STUB: not implemented"; return }
