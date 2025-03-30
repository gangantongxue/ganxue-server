package middleware

import (
	"context"
	"ganxue-server/utils/token"
	"github.com/cloudwego/hertz/pkg/app"
	"strings"
)

func JwtMiddleware() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		// 解析token
		authHeader := ctx.Request.Header.Get("Authorization")
		shortToken := strings.TrimPrefix(authHeader, "Bearer ")

		userID, err := token.ParseToken(shortToken)
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}

		c = context.WithValue(c, "userID", userID)
		ctx.Next(c)
	}
}
