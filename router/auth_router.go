package router

import "ganxue-server/handler"

func AuthRouter() {
	AuthGroup.POST("reset-password", handler.ResetPassword())
	AuthGroup.GET("user/info", handler.UserInfo())
}
