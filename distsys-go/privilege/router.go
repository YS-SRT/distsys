package privilege

import (
	"distsys/global"
	"distsys/global/config"
	"distsys/privilege/controller"
	"distsys/privilege/model"
	"distsys/privilege/service"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PrivilegeModule struct {
	PrivilegeService    *service.PrivilegeService
	PrivilegeController *controller.PrivilegeController
}

func BuildModule(conf *config.ConfigDoc) *PrivilegeModule {

	MigrateSchema()

	e, a := newGlobalCasbinForcer()
	svr := service.BuildService(conf, e, a)
	ctl := controller.BuildController(svr)

	return &PrivilegeModule{

		PrivilegeService:    svr,
		PrivilegeController: ctl,
	}
}

func (p *PrivilegeModule) Routes(base *gin.RouterGroup) {

	usr := base.Group("/privilege")
	{
		usr.GET("/policies", p.PrivilegeController.GetPolicies)
	}
}

func newGlobalCasbinForcer() (*casbin.SyncedEnforcer, *gormadapter.Adapter) {

	gormadapter.TurnOffAutoMigrate(global.GSQLDB)
	if a, err := gormadapter.NewAdapterByDBWithCustomTable(global.GSQLDB, &model.CasbinRule{}); err != nil {

		zap.L().Fatal("setup casbin enforcer of privilege model failed", zap.Error(err))

	} else {

		if e, err := casbin.NewSyncedEnforcer("privilege/conf/casbin.conf", a); err != nil {

			zap.L().Fatal("setup casbin enforcer of privilege model failed", zap.Error(err))

		} else {

			if err := e.LoadPolicy(); err != nil {

				zap.L().Fatal("casbin load policies failed", zap.Error(err))
			}

			// g for user, g2 for resource
			e.AddNamedMatchingFunc("g2", "KeyMatch2", util.KeyMatch2)

			global.GCasbinSF = e
			return e, a

		}

	}

	return nil, nil

}

func MigrateSchema() {

	global.GSQLDB.AutoMigrate(

		&model.CasbinRule{},

		&model.ResourceRole{},
		&model.ResourceRoleMap{},

		&model.UserRole{},
		&model.UserRoleMap{},

		&model.UserResourceGroupMap{},
		&model.RestAction{},
	)

}
