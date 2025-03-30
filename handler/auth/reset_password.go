package auth

import (
	"context"
	"ganxue-server/utils/db/mysql"
	"ganxue-server/utils/password"
	"github.com/cloudwego/hertz/pkg/app"
)

// ResetPassword 重设密码
func ResetPassword() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		user := struct {
			Email       string `json:"email"`
			OldPassword string `json:"old_password"`
			NewPassword string `json:"new_password"`
		}{}
		if err := ctx.Bind(&user); err != nil {
			ctx.JSON(400, map[string]string{"message": "解析请求体失败"})
			return
		}

		// 查找用户
		rawUser, err := mysql.FindUserByEmail(user.Email)
		if err != nil {
			ctx.JSON(400, map[string]string{"message": "用户不存在"})
			return
		}

		// 密码校验
		if !password.ComparePasswords(rawUser.Password, user.OldPassword) {
			ctx.JSON(400, map[string]string{"message": "原密码错误"})
			return
		}

		// 密码加密
		newPw, err := password.EncryptPassword(user.NewPassword)
		if err != nil {
			ctx.JSON(400, map[string]string{"message": "密码加密失败"})
			return
		}

		// 重设密码
		if err := mysql.ChangePassword(rawUser, newPw); err != nil {
			ctx.JSON(500, map[string]string{"message": "服务器错误"})
			return
		}

		ctx.JSON(200, map[string]string{"message": "重设密码成功"})
	}
}
