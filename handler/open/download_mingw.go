package open

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func DownloadMingw() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		zipPath := "./static/download/MinGW.zip"
		filename := "MinGW.zip"

		ctx.Response.Header.Set("Content-Disposition", "attachment; filename="+filename)
		ctx.Response.Header.Set("Content-Type", "application/zip")
		ctx.File(zipPath)
	}
}
