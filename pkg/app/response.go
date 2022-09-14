package app

import (
	"gin_log/pkg/e"
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}
func (g *Gin) Response(httpCode, errorCode int, data any) {
	g.C.JSON(httpCode, gin.H{
		"code": errorCode,
		"msg": e.GetMsg(errorCode),
		"data": data,
	})
}
