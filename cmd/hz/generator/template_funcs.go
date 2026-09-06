package generator

import (
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/cloudwego/hertz/cmd/hz/util"
)

var funcMap = func() template.FuncMap {
	m := template.FuncMap{
		"GetUniqueHandlerOutDir": getUniqueHandlerOutDir,
		"ToSnakeCase":            util.ToSnakeCase,
		"Split":                  strings.Split,
		"Trim":                   strings.Trim,
		"EqualFold":              strings.EqualFold,
	}
	for key, f := range sprig.TxtFuncMap() {
		m[key] = f
	}
	return m
}()

func getUniqueHandlerOutDir(methods []*HttpMethod) (ret []string) {
	_ = "STUB: not implemented"
	return nil
}
