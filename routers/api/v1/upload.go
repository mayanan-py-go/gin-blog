package v1

import (
	"gin_log/pkg/e"
	"gin_log/pkg/logging"
	"gin_log/pkg/upload"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UploadImage(c *gin.Context) {
	code := e.SUCCESS
	data := map[string]any{}

	f, fHeader, err := c.Request.FormFile("image")
	if err != nil {
		logging.Warn(err)
		code = e.ERROR
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg": e.GetMsg(code),
			"data": data,
		})
		return
	}

	if fHeader == nil {
		code = e.INVALID_PARAMS
	}else {
		imageName := upload.GetImageName(fHeader.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(f) {
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		}else{
			err = upload.CheckImage(fullPath)
			if err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
			}else if err = c.SaveUploadedFile(fHeader, src); err != nil {
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			}else{
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}
