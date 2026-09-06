package render

import (
	"html/template"
	"sync"
	"time"

	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/fsnotify/fsnotify"
)

type Delims struct {
	Left string

	Right string
}

type HTMLRender interface {
	Instance(string, interface{}) Render
	Close() error
}

type HTMLProduction struct {
	Template *template.Template
}

type HTML struct {
	Template *template.Template
	Name     string
	Data     interface{}
}

var htmlContentType = "text/html; charset=utf-8"

func (r HTMLProduction) Instance(name string, data interface{}) Render {
	_ = "STUB: not implemented"
	return *new(Render)
}

func (r HTMLProduction) Close() error { _ = "STUB: not implemented"; return nil }

func (r HTML) Render(resp *protocol.Response) error { _ = "STUB: not implemented"; return nil }

func (r HTML) WriteContentType(resp *protocol.Response) { _ = "STUB: not implemented"; return }

type HTMLDebug struct {
	sync.Once
	Template        *template.Template
	RefreshInterval time.Duration

	Files   []string
	FuncMap template.FuncMap
	Delims  Delims

	reloadCh chan struct{}
	watcher  *fsnotify.Watcher
}

func (h *HTMLDebug) Instance(name string, data interface{}) Render {
	_ = "STUB: not implemented"
	return *new(Render)
}

func (h *HTMLDebug) Close() error { _ = "STUB: not implemented"; return nil }

func (h *HTMLDebug) reload() { _ = "STUB: not implemented"; return }

func (h *HTMLDebug) startChecker() { _ = "STUB: not implemented"; return }
