package controllers

import (
	v1 "viki/controllers/api/v1"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func InitRouter() *gin.Engine {

	r := gin.New()

	r.Use(cors.Default())

	r.Use(gin.Logger())
	r.Use(gin.Recovery())


	apiv1 := r.Group("/api/v1")

	{
		//获取标签列表
		apiv1.GET("/articles", v1.GetArticles)
		apiv1.GET("/articles/:id", v1.GetArticle)
		apiv1.POST("/articles", v1.CreateArticle)
		// apiv1.GET("/pods", v1.GetPod)

		// apiv1.GET("/deploys/:name", v1.GetDeploy)
		// apiv1.GET("/deploys", v1.GetDeploys)

	}
	return r
}