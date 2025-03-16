package redis

import (
	"fmt"
	"ganxue-server/global"
	"github.com/go-redis/redis/v8"
	"os"
)

func Init() {
	global.RDB = redis.NewClient(&redis.Options{
		Addr:     global.CONFIG.Redis.Addr,
		Password: global.CONFIG.Redis.Password,
		DB:       global.CONFIG.Redis.DB,
	})
	// 测试连接
	str, _err := global.RDB.Ping(global.CTX).Result()
	if _err != nil {
		fmt.Println("redis连接失败", str, _err)
		os.Exit(1)
	}
}

func Close() {
	err := global.RDB.Close()
	if err != nil {
		fmt.Println("redis关闭失败", err)
		os.Exit(1)
	}
}
