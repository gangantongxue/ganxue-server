package initialize

import (
	"ganxue-server/static"
	"ganxue-server/utils/db"
	"ganxue-server/utils/log"
)

func InitAll() {
	// 加载配置文件
	if err := LoadConfig(); err != nil {
		log.Error(err)
	}
	// 初始化数据库
	db.Init()

	// 初始化日志
	log.Init()

	// 初始化静态文件
	static.Init()
}

func CloseAll() {
	db.Close()
}
