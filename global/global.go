package global

import (
	"context"
	"ganxue-server/config"
	"github.com/fsnotify/fsnotify"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	CONFIG  config.Config
	VIPER   *viper.Viper
	RDB     *redis.Client
	DB      *gorm.DB
	CRON    *cron.Cron
	MDB     *mongo.Database
	MD      *mongo.Collection
	ANSWER  *mongo.Collection
	WATCHER *fsnotify.Watcher
)

var (
	CTX = context.Background()
)
