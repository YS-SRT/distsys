package controller

import (
	"distsys/base"
	"distsys/global/config"
	"distsys/user/convertor"
	"distsys/user/model/dto"
	"distsys/user/service"
	"distsys/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	config  *config.ConfigDoc
	service *service.UserService
}

func BuildController(config *config.ConfigDoc, service *service.UserService) *UserController {
	return &UserController{
		config:  config,
		service: service,
	}
}

// @BasePath /api/v1
// @Summary UserLogin
// @Schemes
// @Description login
// @Tags user
// @Accept json
// @Produce json
// @Param object body dto.UserLoginByPWDInput true "longin params"
// @Success 200 {string} json "{"code":"0", "message":"ok", "data":"response data"}"
// @Router /user/login [post]
func (u *UserController) Login(c *gin.Context) {

	var input *dto.UserLoginByPWDInput
	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, base.BuildBadRequestResp(err.Error()))
		return
	}

	// verify and go check logic
	user, err := u.service.GetUserByPwd(input.Phone, input.Password)
	if err != nil {

		c.JSON(err.Code, base.BuildInternalErrorResp(err.Message))
		return
	}

	token, err := utils.GenerateAccessToken(utils.BuildClaims(user.ID, user.Nickname))
	if err != nil {

		c.JSON(http.StatusInternalServerError, base.BuildInternalErrorResp(err.Error()))
		return
	}

	c.JSON(http.StatusOK, base.BuildSuccessResp(convertor.User2VO(user, token, "")))
}

// @BasePath /api/v1
// @Summary UserRegister
// @Schemes
// @Description Register
// @Tags user
// @Accept json
// @Produce json
// @Param object body dto.UserRegisterInput true "register params"
// @Success 200 {string} json "{"code":"0", "message":"ok", "data":"response data"}"
// @Router /user/register [post]
func (u *UserController) Register(c *gin.Context) {
	var input dto.UserRegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {

		c.JSON(http.StatusBadRequest, base.BuildFailedResp(http.StatusBadRequest, err.Error()))
		return
	}

	user := convertor.RegisterInput2User(&input)
	if user, err := u.service.CreateUser(user); err != nil {

		c.JSON(err.Code, base.BuildFailedResp(err.Code, err.Message))

	} else {

		c.JSON(http.StatusOK, base.BuildSuccessResp(user))
	}

}
