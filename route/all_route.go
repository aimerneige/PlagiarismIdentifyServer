package route

import "github.com/gin-gonic/gin"

func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r = IndexRouteCollection(r)
	r = UserRouteCollection(r)
	r = ArticleRouteCollection(r)

	return r
}
