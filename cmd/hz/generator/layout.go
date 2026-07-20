package generator

type Layout struct {
	OutDir          string
	GoModule        string
	ServiceName     string
	UseApacheThrift bool
	HasIdl          bool
	NeedGoMod       bool
	ModelDir        string
	HandlerDir      string
	RouterDir       string
}

type LayoutGenerator struct {
	ConfigPath string
	TemplateGenerator
}

var (
	layoutConfig  = defaultLayoutConfig
	packageConfig = defaultPkgConfig
)

func SetDefaultTemplateConfig() { _ = "STUB: not implemented"; return }

func (lg *LayoutGenerator) Init() error { _ = "STUB: not implemented"; return nil }

func (lg *LayoutGenerator) checkInited() error { _ = "STUB: not implemented"; return nil }

func (lg *LayoutGenerator) Generate(data map[string]interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (lg *LayoutGenerator) GenerateByService(service Layout) error {
	_ = "STUB: not implemented"
	return nil
}

func serviceToLayoutData(service Layout) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func serviceToRouterData(service Layout) (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (lg *LayoutGenerator) GenerateByConfig(configPath string) error {
	_ = "STUB: not implemented"
	return nil
}

func (lg *LayoutGenerator) Degenerate() error { _ = "STUB: not implemented"; return nil }
