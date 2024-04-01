package middleware

import (
	"distsys/base"
	"distsys/global"
	"distsys/global/adapter"
	"distsys/utils"
	"fmt"
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func VerifyUserToken4Auth() gin.HandlerFunc {

	return func(c *gin.Context) {

		if path := c.Request.URL.Path; strings.HasPrefix(path, "/api/v1") && !slices.Contains(global.GCONFIG.JWT.ExcludePaths, path) {

			tokenString := c.Request.Header.Get("token")
			if tokenString == "" {
				c.AbortWithStatusJSON(http.StatusUnauthorized, base.BuildFailedResp(base.CodeTokenNotTransmitted, "authorizate failed, check token in header"))
				return
			}

			if token, err := utils.ParseAccessToken(tokenString); err != nil {

				c.AbortWithStatusJSON(http.StatusUnauthorized, base.BuildFailedResp(base.CodeTokenNotParsed, err.Error()))
				return

			} else {

				//zap.L().Sugar().Info(token)
				//TODO: check token exist in server side

				path, _ := strings.CutPrefix(path, "/api/v1")
				authId := fmt.Sprintf("%s%s", global.GCONFIG.Casbin.UserPrefix, token["sub"])
				req := &adapter.UserAuthReq{
					UserId:   authId,
					Resource: path,
					Act:      c.Request.Method,
				}
				//zap.L().Sugar().Info(req)
				if ok, err := global.AUserAuth.IsAuthorized(req); err != nil {

					c.AbortWithStatusJSON(http.StatusForbidden, base.BuildFailedResp(http.StatusForbidden, err.Error()))
					return

				} else {

					if !ok {

						c.AbortWithStatusJSON(http.StatusForbidden, base.BuildFailedResp(http.StatusForbidden, "not match rules, access forbidden"))
						return

					}

				}

			}

		}

		c.Next()

	}

}
