package router

import "ganxue-server/handler"

func OpenRouter() {
	OpenGroup.POST("/sign-up", handler.SignUp())
	OpenGroup.POST("/sign-in", handler.SignIn())
	OpenGroup.POST("/reset-password", handler.ResetPassword())
	OpenGroup.POST("/refresh", handler.Refresh())

	OpenGroup.GET("/ver-code", handler.VerCode())
	OpenGroup.GET("/delete-user", handler.DeleteUser())
}
