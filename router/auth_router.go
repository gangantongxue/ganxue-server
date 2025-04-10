package router

import (
	"ganxue-server/handler/auth"
)

func AuthRouter() {
	AuthGroup.POST("/reset-password", auth.ResetPassword())
	AuthGroup.POST("/run-code", auth.RunCode())

	AuthGroup.GET("/user/info", auth.UserInfo())
	AuthGroup.GET("/get-docs", auth.GetDocs())
	AuthGroup.GET("/detailed-user-info", auth.DetailedUserInfo())
	AuthGroup.GET("/get-catalogue", auth.GetCatalogue())
}
