package route

import (
	"restful-template/controllers"
	"restful-template/middleware"

	"github.com/gin-gonic/gin"
)

func ArticleRouteCollection(r *gin.Engine) *gin.Engine {
	article := r.Group("/article")
	article.Use(middleware.AuthMiddleware())
	article.GET("", controllers.GetArticle)
	article.DELETE("", controllers.DeleteArticle)
	article.Use(middleware.ArticleMiddleware())
	article.POST("", controllers.CreateArticle)
	article.PUT("", controllers.UpdateArticle)

	return r
}
