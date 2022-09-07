package routers

import (
	"gin_log/pkg/setting"
	v1 "gin_log/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "test"})
	})

	apiv1 := r.Group("/api/v1")
	// tag
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新指定标签
		apiv1.PUT("/tag/:id", v1.EditTag)
		// 删除指定标签
		apiv1.DELETE("/tag/:id", v1.DeleteTag)
	}

	return r
}
