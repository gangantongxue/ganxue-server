package db

import (
	"ganxue-server/utils/db/mysql"
	"ganxue-server/utils/db/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
