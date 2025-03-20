package auth

import (
	"context"
	"ganxue-server/utils/db/mysql"
	"github.com/cloudwego/hertz/pkg/app"
)

func DetailedUserInfo() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		userID := c.Value("userID")
		if userID == nil {
			ctx.AbortWithStatus(401)
			return
		}
		user, err := mysql.FindUserByID(userID.(uint))
		if err != nil {
			ctx.AbortWithStatus(401)
			return
		}
		ctx.JSON(200, map[string]string{
			"user_name": user.UserName,
			"email":     user.Email,
			"create_at": user.CreatedAt.String(),
		})
	}
}
