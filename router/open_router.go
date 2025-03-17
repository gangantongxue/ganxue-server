package router

import (
	"ganxue-server/handler/open"
)

func OpenRouter() {
	OpenGroup.POST("/sign-up", open.SignUp())
	OpenGroup.POST("/sign-in", open.SignIn())
	OpenGroup.POST("/forget-password", open.ForgetPassword())

	OpenGroup.GET("/refresh", open.Refresh())
	OpenGroup.GET("/ver-code", open.VerCode())
	OpenGroup.GET("/delete-user", open.DeleteUser())
}
