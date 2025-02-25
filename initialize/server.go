package initialize

import (
	"ganxue-server/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
)

// Routers 初始化总路由
func Routers() *server.Hertz {
	h := server.New(server.WithHostPorts(":8080"))

	h.Use(middleware.LogMid())

	return h
}
