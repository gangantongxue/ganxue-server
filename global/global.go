package global

import (
	"context"
	"ganxue-server/config"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	CONFIG config.Config
	VIPER  *viper.Viper
	RDB    *redis.Client
	DB     *gorm.DB
	CRON   *cron.Cron
	MDB    *mongo.Database
)

var (
	CTX = context.Background()
)
