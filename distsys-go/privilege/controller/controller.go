package controller

import (
	"distsys/base"
	"distsys/privilege/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PrivilegeController struct {
	PrivilegeService *service.PrivilegeService
}

func BuildController(svr *service.PrivilegeService) *PrivilegeController {

	return &PrivilegeController{

		PrivilegeService: svr,
	}

}

func (p *PrivilegeController) GetPolicies(c *gin.Context) {

	c.JSON(http.StatusOK, base.BuildSuccessResp(p.PrivilegeService.SyncedEnforcer.GetPolicy()))

}

func (p *PrivilegeController) GetUserRoles(c *gin.Context) {

	c.JSON(http.StatusOK, base.BuildSuccessResp(p.PrivilegeService.SyncedEnforcer.GetAllNamedRoles(p.PrivilegeService.Config.Casbin.UserGroupPType)))
}

func (p *PrivilegeController) GetResourceRoles(c *gin.Context) {

	c.JSON(http.StatusOK, base.BuildSuccessResp(p.PrivilegeService.SyncedEnforcer.GetAllNamedRoles(p.PrivilegeService.Config.Casbin.ResourceGroupPType)))
}
