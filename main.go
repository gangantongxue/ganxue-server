package main

// TODO:ver-code

import (
	"ganxue-server/initialize"
	"ganxue-server/router"
)

func main() {
	// 初始化
	initialize.InitAll()

	// 初始化路由
	h := initialize.Routers()
	router.Init(h)

	// 启动服务
	h.Spin()
}
