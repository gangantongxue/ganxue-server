package handler

import (
	"context"
	"ganxue-server/utils/token"
	"github.com/cloudwego/hertz/pkg/app"
)

// Refresh 刷新短token
func Refresh() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		longToken := ctx.Cookie("long_token")
		userID, err := token.ParseToken(string(longToken))
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}
		shortToken, err := token.GenerateShortToken(userID)
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}
		ctx.JSON(200, map[string]string{
			"token": shortToken,
		})
		return
	}
}
