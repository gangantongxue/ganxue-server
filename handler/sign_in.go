package handler

import (
	"context"
	"ganxue-server/model/user_model"
	"ganxue-server/utils/db/mysql"
	"ganxue-server/utils/log"
	"ganxue-server/utils/password"
	"ganxue-server/utils/token"
	"github.com/cloudwego/hertz/pkg/app"
)

// SignIn 登录
func SignIn() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		user := user_model.User{}
		// 解析请求体
		if err := ctx.Bind(&user); err != nil {
			ctx.String(400, "请求错误")
			return
		}
		// 查找用户
		_user, err := mysql.FindUserByEmail(user.Email)
		// 用户不存在
		if err != nil {
			ctx.String(400, "未注册")
			return
		}

		// 密码校验
		if !password.ComparePasswords(_user.Password, user.Password) {
			ctx.String(400, "密码错误")
			return
		}

		// 生成token
		var shortToken, longToken string
		if shortToken, err = token.GenerateShortToken(user.ID); err != nil {
			log.Error(err)
			ctx.String(500, "服务器错误")
			return
		}
		if longToken, err = token.GenerateLongToken(user.ID); err != nil {
			log.Error(err)
			ctx.String(500, "服务器错误")
			return
		}
		ctx.JSON(200, struct {
			Message string `json:"message"`
			Data    struct {
				Email      string `json:"email"`
				UserName   string `json:"user_name"`
				ShortToken string `json:"short_token"`
				LongToken  string `json:"long_token"`
			} `json:"data"`
		}{
			Message: "注册成功",
			Data: struct {
				Email      string `json:"email"`
				UserName   string `json:"user_name"`
				ShortToken string `json:"short_token"`
				LongToken  string `json:"long_token"`
			}{
				Email:      user.Email,
				UserName:   user.UserName,
				ShortToken: shortToken,
				LongToken:  longToken,
			},
		})
	}
}
