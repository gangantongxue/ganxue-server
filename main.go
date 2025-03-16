package main

import (
	"ganxue-server/initialize"
	"ganxue-server/router"
	_ "net/http/pprof"
)

func main() {
	// 初始化
	initialize.InitAll()
	defer initialize.CloseAll()

	// 初始化路由
	h := initialize.Routers()
	router.Init(h)

	// 启动服务
	h.Spin()
}
