package db

import (
	"ganxue-server/utils/db/mongodb"
	"ganxue-server/utils/db/mysql"
	"ganxue-server/utils/db/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
	mongodb.Init()
}

func Close() {
	mysql.Close()
	redis.Close()
	mongodb.Close()
}
