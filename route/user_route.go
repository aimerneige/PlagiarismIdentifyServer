package route

import (
	"restful-template/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouteCollection(r *gin.Engine) *gin.Engine {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	return r
}
