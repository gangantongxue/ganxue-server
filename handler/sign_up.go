package handler

import (
	"context"
	"ganxue-server/model/user_model"
	"ganxue-server/utils/db/mysql"
	"ganxue-server/utils/log"
	"ganxue-server/utils/password"
	"ganxue-server/utils/token"
	"ganxue-server/utils/ver_code"
	"github.com/cloudwego/hertz/pkg/app"
)

// SignUp 创建用户
// OpenAPI
// @Summary 创建用户
// @Description 创建用户
// @Tags 用户
// @Accept json
// @Produce json
// @Param user body user_model.User true "用户信息"
// @Success 200 {object} user_model.User
// @Failure 400 {object} error.Error
// @Failure 500 {object} error.Error
// @Router /log-up [post]
func SignUp() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		user := struct {
			user_model.User
			VerCode string `json:"ver_code"`
		}{}
		// 解析请求体
		if err := ctx.Bind(&user); err != nil {
			ctx.String(400, "请求错误")
			return
		}
		// 验证验证码
		if !ver_code.Verify(user.Email, user.VerCode) {
			ctx.String(400, "验证码错误")
			return
		}
		// 查找用户
		_user, err := mysql.FindUserByEmail(user.Email)
		// 用户已存在
		if err == nil {
			ctx.JSON(400, struct {
				Message string `json:"message"`
				Data    struct {
					Email    string `json:"email"`
					UserName string `json:"user_name"`
				} `json:"data"`
			}{
				Message: "用户已存在",
				Data: struct {
					Email    string `json:"email"`
					UserName string `json:"user_name"`
				}{
					Email:    _user.Email,
					UserName: _user.UserName,
				},
			})
			return
		}

		// 密码加密
		if pwd, err := password.EncryptPassword(user.Password); err != nil {
			log.Error(err)
			ctx.String(500, "服务器错误")
			return
		} else {
			user.Password = pwd
		}

		// 创建用户
		if err := mysql.Create(&user); err != nil {
			log.Error(err)
			ctx.String(500, "服务器错误")
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
		return
	}
}
