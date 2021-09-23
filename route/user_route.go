package route

import (
	"plagiarism-identify-server/controllers"

	"github.com/gin-gonic/gin"
)

func UserRouteCollection(r *gin.Engine) *gin.Engine {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	return r
}
