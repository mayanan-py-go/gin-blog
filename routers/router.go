package routers

import (
	"gin_log/middleware/jwt"
	"gin_log/pkg/export"
	"gin_log/pkg/setting"
	"gin_log/pkg/upload"
	v1 "gin_log/routers/api/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	r.GET("/test", func(context *gin.Context) {
		context.JSON(200, gin.H{"message": "test"})
	})

	r.GET("/auth", v1.GetAuth)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))

	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())

	{
		apiv1.POST("/upload", v1.UploadImage)
		apiv1.POST("/tags/export", v1.ExportTag)
		apiv1.POST("/tags/import", v1.ImportTag)
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

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

	// article
	{
		// 获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		// 获取指定文章
		apiv1.GET("/article/:id", v1.GetArticle)
		// 新建文章
		apiv1.POST("/article", v1.AddArticle)
		// 更新指定文章
		apiv1.PUT("/article/:id", v1.EditArticle)
		// 删除指定文章
		apiv1.DELETE("/article/:id", v1.DeleteArticle)
	}

	return r
}
