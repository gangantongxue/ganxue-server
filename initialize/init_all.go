package initialize

import (
	"ganxue-server/router"
	"ganxue-server/utils/db"
	"ganxue-server/utils/log"
)

func InitAll() {
	// 加载配置文件
	if err := LoadConfig(); err != nil {
		log.Error(err)
	}

	// 初始化日志
	log.Init()

	// 初始化数据库
	db.Init()

	// 初始化路由
	h := Routers()
	router.Init(h)

	// 启动服务
	h.Spin()
}
