package router

import (
	"ganxue-server/handler/auth"
)

func AuthRouter() {
	AuthGroup.POST("reset-password", auth.ResetPassword())
	AuthGroup.GET("user/info", auth.UserInfo())
	AuthGroup.GET("/get-docs", auth.GetDocs())
}
