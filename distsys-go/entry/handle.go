package entry

import (
	"distsys/base"
	"distsys/global"
	"distsys/utils"
	"errors"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (entry *Entry) home(c *gin.Context) {

	c.HTML(http.StatusOK, "login.html", gin.H{})
}

func (entry *Entry) navigate(c *gin.Context) {

	c.HTML(http.StatusOK, "navigation.html", gin.H{})
}

func (entry *Entry) foot(c *gin.Context) {

	c.HTML(http.StatusOK, "foot.html", gin.H{})
}

// @BasePath /api/v1
// @Summary UploadFile
// @Schemes
// @Description upload file
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "file"
// @Success 200 {string} json "{"code":"0", "message":"ok", "data":"response data"}"
// @Router /public/upload [post]
func (entry *Entry) uploadFile(c *gin.Context) {

	file, _ := c.FormFile("file")
	fileObj, err := file.Open()

	if err == nil {

		fileName := strings.ReplaceAll(uuid.NewString(), "-", "") + path.Ext(file.Filename)

		if ok := utils.UploadFile(fileName, fileObj, file.Size); ok {

			//TODO Store Path into DB
			c.JSON(http.StatusOK, base.BuildSuccessResp(gin.H{"path": global.GCONFIG.MinioAuth.PublicPrefix + fileName}))
			return

		} else {

			err = errors.New("internal service error, can't put object into minio")
		}

	}

	defer fileObj.Close()

	zap.L().Error("upload file failed: ", zap.Error(err))

	c.JSON(http.StatusOK, base.BuildFailedResp(http.StatusBadRequest, "upload File failed"))

}

// @BasePath /api/v1
// @Summary UploadFiles
// @Schemes
// @Description upload Files
// @Tags user
// @Accept multipart/form-data
// @Produce json
// @param file formData []file false "files to upload"
// @Success 200 {string} json "{"code":"0", "message":"ok", "data":"response data"}"
// @Router /public/uploads [post]
func (entry *Entry) uploadFiles(c *gin.Context) {

	err := c.Request.ParseMultipartForm(8 << 20) // 8Mb
	if err != nil {

		c.JSON(http.StatusBadRequest, base.BuildFailedResp(http.StatusBadRequest, "File is more than 8MB"))
	}
	formdata := c.Request.MultipartForm
	files := formdata.File["file"]
	var uploadedFiles []string

	for _, v := range files {

		fileObj, err := v.Open()

		if err == nil {

			fileName := strings.ReplaceAll(uuid.NewString(), "-", "") + path.Ext(v.Filename)

			if ok := utils.UploadFile(fileName, fileObj, v.Size); !ok {

				err = errors.New("internal service error, can't put object into minio")

			} else {

				uploadedFiles = append(uploadedFiles, global.GCONFIG.MinioAuth.PublicPrefix+fileName)
				continue
			}
		}

		defer fileObj.Close()

		zap.L().Error("upload files failed", zap.Error(err))

		c.JSON(http.StatusBadRequest, base.BuildFailedResp(http.StatusBadRequest, "upload File failed"))
		return
	}

	c.JSON(http.StatusOK, base.BuildSuccessResp(gin.H{"paths": uploadedFiles}))
}
