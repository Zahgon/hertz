package generator

type File struct {
	Path        string
	Content     string
	NoRepeat    bool
	FileTplName string
}

func (file *File) Lint() error { _ = "STUB: not implemented"; return nil }
