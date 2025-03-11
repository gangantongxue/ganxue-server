package router

import "ganxue-server/handler"

func OpenRouter() {
	OpenGroup.POST("/sign-up", handler.SignUp())
	OpenGroup.POST("/sign-in", handler.SignIn())
	OpenGroup.POST("/refresh", handler.Refresh())
	OpenGroup.POST("/forget-password", handler.ForgetPassword())

	OpenGroup.GET("/ver-code", handler.VerCode())
	OpenGroup.GET("/delete-user", handler.DeleteUser())
}
