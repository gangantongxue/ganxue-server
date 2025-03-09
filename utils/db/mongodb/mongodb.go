package mongodb

import (
	"fmt"
	"ganxue-server/global"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func Init() {
	// 连接MongoDB
	uri := fmt.Sprintf("mongodb://%s:%s", global.CONFIG.Mongo.Host, global.CONFIG.Mongo.Port)
	client, err := mongo.Connect(global.CTX, options.Client().ApplyURI(uri))
	if err != nil {
		fmt.Println("MongoDB连接失败", err)
		os.Exit(1)
	}

	// 测试连接
	_err := client.Ping(global.CTX, nil)
	if _err != nil {
		fmt.Println("MongoDB连接失败", _err)
		os.Exit(1)
	}
	global.MDB = client.Database(global.CONFIG.Mongo.Database)

}
