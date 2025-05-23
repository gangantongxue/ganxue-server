package global

import (
	"context"
	"ganxue-server/config"
	"ganxue-server/model/catalogue_model"
	"github.com/go-redis/redis/v8"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"
)

var (
	CONFIG         config.Config
	VIPER          *viper.Viper
	RDB            *redis.Client
	DB             *gorm.DB
	CRON           *cron.Cron
	MDB            *mongo.Database
	MD             *mongo.Collection
	ANSWER         *mongo.Collection
	GROUP          singleflight.Group
	RUN_CODE_MUTEX sync.Mutex
	CATALOGUE      catalogue_model.CatalogueModel
)

var (
	CTX = context.Background()
)
