package route

import (
	"restful-template/controllers"

	"github.com/gin-gonic/gin"
)

func IndexRouteCollection(r *gin.Engine) *gin.Engine {
	r.GET("/", controllers.Index)

	return r
}
