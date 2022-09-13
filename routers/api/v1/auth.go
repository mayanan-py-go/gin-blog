package v1

import (
	"gin_log/models"
	"gin_log/pkg/e"
	"gin_log/pkg/logging"
	"gin_log/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
}
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	valid := validation.Validation{}
	au := auth{Username: username, Password: password}
	ok, _ := valid.Valid(&au)

	code := e.INVALID_PARAMS
	data := map[string]any{}

	if ok {
		exist := models.CheckAuth(username, password)
		if exist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			}else{
				data["token"] = token
				code = e.SUCCESS
			}
		}else{
			code = e.ERROR_AUTH
		}
	}else{
		for _, err := range valid.Errors {
			logging.Error(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg": e.GetMsg(code),
		"data": data,
	})
}











