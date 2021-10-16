// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func HomeworkRouteCollection(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	homework := v1.Group("/homework")

	homeworkInfoGet := homework.Group("")
	homeworkInfoGet.Use(middleware.UserAuthMiddleware())
	homeworkInfoGet.GET("", controllers.HomeworkInfoGet)

	return r
}
