package auth

import (
	"context"
	"ganxue-server/global"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetCatalogue() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		ctx.JSON(200, global.CATALOGUE)
	}
}
