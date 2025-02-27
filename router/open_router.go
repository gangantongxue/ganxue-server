package router

import "ganxue-server/handler"

func OpenRouter() {
	OpenGroup.POST("/sign-up", handler.SignUp())
	OpenGroup.POST("/sign-in", handler.SignIn())
	OpenGroup.GET("/ver-code", handler.VerCode())
}
