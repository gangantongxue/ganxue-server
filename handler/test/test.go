package test

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
)

func Test() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.String(200, "succeed")
	}
}
