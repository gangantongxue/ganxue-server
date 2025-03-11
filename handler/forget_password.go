package handler

import (
	"context"
	"ganxue-server/utils/db/mysql"
	"ganxue-server/utils/password"
	"ganxue-server/utils/ver_code"
	"github.com/cloudwego/hertz/pkg/app"
)

func ForgetPassword() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		user := struct {
			Email       string `json:"email"`
			NewPassword string `json:"new_password"`
			VerCode     string `json:"ver_code"`
		}{}
		if err := ctx.Bind(&user); err != nil {
			ctx.JSON(400, map[string]string{"message": "解析请求体失败"})
			return
		}

		// 验证验证码
		if !ver_code.Verify(user.Email, user.VerCode) {
			ctx.JSON(400, map[string]string{"message": "验证码错误"})
			return
		}

		// 查找用户
		rawUser, err := mysql.FindUserByEmail(user.Email)
		if err != nil {
			ctx.JSON(400, map[string]string{"message": "用户不存在"})
			return
		}

		// 密码加密
		newPw, err := password.EncryptPassword(user.NewPassword)
		if err != nil {
			ctx.JSON(500, map[string]string{"message": "密码加密失败"})
			return
		}

		if err := mysql.ChangePassword(rawUser, newPw); err != nil {
			ctx.JSON(500, map[string]string{"message": "服务器错误"})
			return
		}

		ctx.JSON(200, map[string]string{"message": "重设密码成功"})
	}
}
