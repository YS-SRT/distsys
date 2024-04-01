package global

import (
	"distsys/global/adapter"
	"distsys/global/config"

	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// components, , need setup in main func
var (
	GViper    *viper.Viper
	GCONFIG   *config.ConfigDoc
	GEngine   *gin.Engine
	GSQLDB    *gorm.DB
	GLogger   *zap.Logger
	GMinioCli *minio.Client
	GRedisCli *redis.Client

	GCasbinSF *casbin.SyncedEnforcer
)

// adapter interface, need setup in main func
var (
	AUserInfo adapter.UserInfoAdapter
	AUserAuth adapter.UserAuthorizedAdapter
)
