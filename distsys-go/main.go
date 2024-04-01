package main

import (
	"context"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"distsys/entry"
	"distsys/global"
	"distsys/middleware"
	"distsys/privilege"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"moul.io/zapgorm2"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"

	usr "distsys/user"
)

func main() {

	setupConfig()
	setupZap()
	setupSQLDB()
	setupRedis()
	setupMinio()
	setupSever()
}

func setupConfig() {
	configFile := *flag.String("c", "application", "load config file with input name")
	flag.Parse()
	v := getViper(configFile)
	if err := v.Unmarshal(&global.GCONFIG); err != nil {

		log.Fatal(err)
		return
	}
}

func setupZap() {

	entry.SetupZap()
}

func setupSQLDB() {

	if db, err := gorm.Open(postgres.Open(global.GCONFIG.Postgres.DSN), &gorm.Config{Logger: zapgorm2.New(zap.L())}); err != nil {

		zap.L().Fatal("open data base error: ", zap.Error(err))
		return

	} else {

		if pg, err := db.DB(); err != nil {

			zap.L().Fatal("try to get sql.DB error: ", zap.Error(err))
			return

		} else {

			pg.SetMaxIdleConns(global.GCONFIG.Postgres.MaxIdleConns)
			pg.SetMaxOpenConns(global.GCONFIG.Postgres.MaxOpenConns)
			global.GSQLDB = db
		}

	}
}

func setupMinio() {

	cli, err := minio.New(global.GCONFIG.MinioAuth.EndPoint, &minio.Options{
		Creds:  credentials.NewStaticV4(global.GCONFIG.MinioAuth.AccessKey, global.GCONFIG.MinioAuth.AccessSecret, ""),
		Secure: global.GCONFIG.MinioAuth.UseSSL,
	})
	if err != nil {
		zap.L().Fatal("minio client setup failed, err %+v", zap.Error(err))
	}

	global.GMinioCli = cli
}

func setupRedis() {

	if opt, err := redis.ParseURL(global.GCONFIG.Redis.DSN); err != nil {

		zap.L().Fatal("redis client setup failed err %+v", zap.Error(err))

	} else {

		global.GRedisCli = redis.NewClient(opt)
	}
}

func setupSever() {
	//setup engine
	engine := gin.New()
	engine.Use(middleware.GinLogger(), middleware.GinRecovery(false), middleware.Cors())
	engine.Use(middleware.VerifyUserToken4Auth())
	//engine.Use(middleware.LoadTls())
	gin.SetMode(global.GCONFIG.System.Mode)
	engine.MaxMultipartMemory = 8 << 20 //8M
	global.GEngine = engine
	//setup modulea with router
	enter := &entry.Entry{
		UserModule:      usr.BuildModule(),
		PrivilegeModule: privilege.BuildModule(global.GCONFIG),
	}
	enter.Routes()

	//setup global adapter
	global.AUserInfo = enter.UserModule.UserService
	global.AUserAuth = enter.PrivilegeModule.PrivilegeService

	//setup web server
	srv := &http.Server{
		Addr:    global.GCONFIG.System.Addr,
		Handler: engine,
	}
	//run in goruntine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Fatal("listen: %s\n", zap.Error(err))
		}
	}()

	//for gracefull shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	zap.L().Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {

		zap.L().Fatal("Server Shutdown: ", zap.Error(err))
	}
	zap.L().Info("Server exiting")

}

func getViper(fileName string) *viper.Viper {

	config := viper.New()
	config.SetConfigName(fileName)
	config.AddConfigPath(".")
	if err := config.ReadInConfig(); err != nil {
		//zap not set in this time
		log.Fatal("Loading Wrong Config File!", err)
	}
	return config
}
