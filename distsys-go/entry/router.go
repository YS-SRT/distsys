package entry

import (
	"distsys/docs"
	"distsys/global"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (entry *Entry) Routes() {

	//static folds
	global.GEngine.StaticFS(global.GCONFIG.System.StaticPath, http.Dir(global.GCONFIG.System.StaticPath))

	//admin UI
	global.GEngine.LoadHTMLGlob("./admin/pages/*")
	global.GEngine.GET("/", entry.home)
	global.GEngine.GET("/navigation.html", entry.navigate)
	global.GEngine.GET("/foot.html", entry.foot)

	//health check
	global.GEngine.GET("/health", func(c *gin.Context) { c.JSON(200, "ok") })

	//swagger UI
	docs.SwaggerInfo.BasePath = "/api/v1"
	global.GEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//API groups
	v1 := global.GEngine.Group("/api/v1")
	entry.sharedApi(v1)
	entry.UserModule.Routes(v1)
	entry.PrivilegeModule.Routes(v1)

}

func (entry *Entry) sharedApi(base *gin.RouterGroup) {

	pub := base.Group("/public")
	{
		pub.POST("/upload", entry.uploadFile)
		pub.POST("/uploads", entry.uploadFiles)
	}
}
