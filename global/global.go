package global

import (
	"context"
	"ganxue-server/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	CONFIG config.Config
	VIPER  *viper.Viper
	RDB    *redis.Client
	DB     *gorm.DB
)

var (
	CTX = context.Background()
)
