package main

import (
	"context"
	"html/template"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func formatAsDate(t time.Time) string { _ = "STUB: not implemented"; return "" }

func main() {

	h := server.Default(server.WithAutoReloadRender(true, 0))

	h.Delims("{[{", "}]}")

	h.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})
	h.LoadHTMLGlob("./examples/html_rendering/*")

	h.GET("/index", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "index.tmpl", utils.H{
			"title": "Main website",
		})
	})

	h.GET("/raw", func(c context.Context, ctx *app.RequestContext) {
		ctx.HTML(consts.StatusOK, "template.html", utils.H{
			"now": time.Date(2017, 0o7, 0o1, 0, 0, 0, 0, time.UTC),
		})
	})

	h.Spin()
}
