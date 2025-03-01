package middleware

import (
	"context"
	"ganxue-server/utils/token"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
)

func JwtMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 解析cookie
		shortToken := ctx.Cookie("shortToken")

		userID, err := token.ParseToken(string(shortToken))
		if err != nil {
			longToken := ctx.Cookie("longToken")
			userID, err = token.ParseToken(string(longToken))
			if err != nil {
				ctx.AbortWithStatus(401)
				return
			}
			newShortToken, err := token.GenerateShortToken(userID)
			if err != nil {
				ctx.AbortWithStatus(401)
				return
			}
			ctx.SetCookie(
				"shortToken",
				newShortToken,
				60*15,
				"/",
				"",
				protocol.CookieSameSiteDefaultMode,
				false,
				true)
		}
		c = context.WithValue(c, "userID", userID)
		ctx.Next(c)
	}
}
