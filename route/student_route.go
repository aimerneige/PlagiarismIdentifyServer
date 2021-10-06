// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func StudentRouteCollection(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	student := v1.Group("/student")

	studentRegister := student.Group("")
	studentRegister.Use(middleware.RegisterMiddleware())
	studentRegister.POST("", controllers.StudentRegister)

	return r
}
