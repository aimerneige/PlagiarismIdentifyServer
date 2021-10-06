// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRouteCollection(r *gin.Engine) *gin.Engine {
	v1 := r.Group("/api/v1")
	login := v1.Group("/login")

	login.GET("/teacher", controllers.TeacherLogin)
	login.GET("/student", controllers.StudentLogin)

	return r
}
