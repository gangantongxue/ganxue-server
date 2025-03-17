package mongodb

import (
	"fmt"
	"ganxue-server/global"
	"ganxue-server/utils/error"
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
	global.MD = global.MDB.Collection("markdown")
	global.ANSWER = global.MDB.Collection("answer")
}

func Close() {
	err := global.MDB.Client().Disconnect(global.CTX)
	if err != nil {
		return
	}
}

// Find 查找数据
func Find(collection *mongo.Collection, filter interface{}, result interface{}) *error.Error {
	err := collection.FindOne(global.CTX, filter).Decode(&result)
	if err != nil {
		return error.New(error.MongoError, err, "MongoDB查询失败")
	}
	return nil
}

// Insert 插入数据
func Insert(collection *mongo.Collection, data interface{}) *error.Error {
	_, err := collection.InsertOne(global.CTX, data)
	if err != nil {
		return error.New(error.MongoError, err, "MongoDB插入失败")
	}
	return nil
}

// Update 更新数据
func Update(collection *mongo.Collection, filter interface{}, update interface{}, upsert bool) *error.Error {
	opts := options.Update().SetUpsert(upsert)
	_, err := collection.UpdateOne(global.CTX, filter, update, opts)
	if err != nil {
		return error.New(error.MongoError, err, "MongoDB更新失败")
	}
	return nil
}

// Delete 删除数据
func Delete(collection *mongo.Collection, filter interface{}) *error.Error {
	_, err := collection.DeleteOne(global.CTX, filter)
	if err != nil {
		return error.New(error.MongoError, err, "MongoDB删除失败")
	}
	return nil
}
