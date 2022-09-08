package jwt

import (
	"gin_log/pkg/e"
	"gin_log/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func JWT() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			code int
			data any
		)
		code = e.SUCCESS
		token := context.GetHeader("token")
		if token == "" {
			code = e.INVALID_PARAMS
		}else{
			_, err := util.ParseToken(token)
			if err != nil {
				// 1. token过期
				if strings.HasPrefix(err.Error(), "token is expired by") {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				// 2. token解析失败
				}else {
					code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				}
			}
		}
		if code != e.SUCCESS {
			context.JSON(http.StatusOK, gin.H{
				"code": code,
				"msg": e.GetMsg(code),
				"data": data,
			})
			context.Abort()
			return
		}
		context.Next()
	}
}









