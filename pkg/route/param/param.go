package param

type Param struct {
	Key   string
	Value string
}

type Params []Param

func (ps Params) Get(name string) (string, bool) { _ = "STUB: not implemented"; return "", false }

func (ps Params) ByName(name string) (va string) { _ = "STUB: not implemented"; return "" }
