package global

import (
	"context"
	"ganxue-server/config"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
)

var (
	CONFIG    config.Config
	VIPER     *viper.Viper
	RDB       *redis.Client
	DB        *gorm.DB
	CRON      *cron.Cron
	MDB       *mongo.Database
	GO_MD     *mongo.Collection
	GO_ANSWER *mongo.Collection
	GROUP     singleflight.Group
)

var (
	CTX = context.Background()
)
