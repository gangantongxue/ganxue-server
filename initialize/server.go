package initialize

import (
	"ganxue-server/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/cors"
	"net/http"
)

// Routers 初始化总路由
func Routers() *server.Hertz {
	h := server.New(server.WithHostPorts(":8080"))

	c := cors.New(cors.Config{
		AllowOrigins:     []string{"http://127.0.0.1:5500", "http://115.190.92.245:80", "http://127.0.0.1:80", "http://115.190.92.245:443", "https://115.190.92.245:443"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	})

	h.Use(c)

	h.Use(middleware.LogMid())

	return h
}
