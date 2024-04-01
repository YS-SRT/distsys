package user

import (
	"distsys/global"
	"distsys/user/controller"
	"distsys/user/model"
	"distsys/user/service"

	"github.com/gin-gonic/gin"
)

type UserModule struct {
	UserService    *service.UserService
	UserController *controller.UserController
}

func BuildModule() *UserModule {

	svr := service.BuildService(global.GCONFIG, global.GSQLDB)
	ctl := controller.BuildController(global.GCONFIG, svr)

	//init data table
	MigrateSchema()

	return &UserModule{

		UserService:    svr,
		UserController: ctl,
	}

}

func (m *UserModule) Routes(base *gin.RouterGroup) {

	usr := base.Group("/user")
	{

		usr.POST("/login", m.UserController.Login)
		usr.POST("/register", m.UserController.Register)
	}
}

func MigrateSchema() {

	global.GSQLDB.AutoMigrate(

		&model.User{},
		&model.Address{},

		&model.Score{},
		&model.Level{},

		&model.Group{},
		&model.Member{},
	)

}
