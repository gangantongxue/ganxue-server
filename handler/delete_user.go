package handler

import (
	"context"
	"ganxue-server/utils/db/mysql"
	"github.com/cloudwego/hertz/pkg/app"
)

// DeleteUser 删除用户
func DeleteUser() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		email := ctx.Query("email")
		if email == "" {
			ctx.String(400, "邮箱不能为空")
			return
		}

		user, err := mysql.FindUserByEmail(email)
		if err != nil {
			ctx.String(500, "服务器错误")
			return
		}
		if user == nil {
			ctx.String(400, "用户不存在")
			return
		}

		if err := mysql.Delete(user); err != nil {
			ctx.String(500, "服务器错误")
			return
		}

		ctx.String(200, "删除用户成功")
		return
	}
}
