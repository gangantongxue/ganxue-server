package redis

import (
	"ganxue-server/global"
	"ganxue-server/utils/log"
	"github.com/go-redis/redis/v8"
)

func Init() {
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     global.CONFIG.Redis.Addr,
		Password: global.CONFIG.Redis.Password,
		DB:       global.CONFIG.Redis.DB,
	})
	// 测试连接
	str, _err := global.RDB.Ping(global.CTX).Result()
	log.Fatal(_err, "连接Redis失败", str)
}
