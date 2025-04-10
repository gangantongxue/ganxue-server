package router

import (
	"ganxue-server/middleware"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/route"
)

var (
	AuthGroup *route.RouterGroup
	OpenGroup *route.RouterGroup
)

func Init(h *server.Hertz) {
	AuthGroup = h.Group("/api/auth", middleware.JwtMiddleware())
	OpenGroup = h.Group("/api/open")

	OpenRouter()
	AuthRouter()
}
