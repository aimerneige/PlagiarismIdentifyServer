// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func TeacherRouteCollection(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	teacher := v1.Group("/teacher")

	teacherRegister := teacher.Group("")
	teacherRegister.Use(middleware.RegisterMiddleware())
	teacherRegister.POST("", controllers.TeacherRegister)

	return r
}
