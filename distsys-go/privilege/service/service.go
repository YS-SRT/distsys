package service

import (
	"distsys/base"
	"distsys/global/adapter"
	"distsys/global/config"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
)

type PrivilegeService struct {
	Config         *config.ConfigDoc
	SyncedEnforcer *casbin.SyncedEnforcer
	GormAdapter    *gormadapter.Adapter
}

func BuildService(conf *config.ConfigDoc, forcer *casbin.SyncedEnforcer, adapter *gormadapter.Adapter) *PrivilegeService {

	return &PrivilegeService{
		Config:         conf,
		SyncedEnforcer: forcer,
		GormAdapter:    adapter,
	}
}

func (p *PrivilegeService) IsAuthorized(req *adapter.UserAuthReq) (bool, *base.CustomError) {

	if ok, err := p.SyncedEnforcer.Enforce(req.UserId, req.Resource, req.Act); err != nil {

		zap.L().Error("check casbin rules failed", zap.Error(err))
		return false, base.BuildCustomForbiddenError(err.Error())

	} else {

		return ok, nil
	}
}

func (p *PrivilegeService) RefreshCasin() error {

	return p.SyncedEnforcer.LoadPolicyFast()
}
