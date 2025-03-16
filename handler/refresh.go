package handler

import (
	"context"
	"ganxue-server/utils/log"
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
			log.Debug("Parse token error")
			return
		}
		shortToken, err := token.GenerateShortToken(userID)
		if err != nil {
			ctx.AbortWithStatus(401)
			log.Debug("Generate short token error")
			return
		}
		ctx.JSON(200, map[string]string{
			"token": shortToken,
		})
		return
	}
}
