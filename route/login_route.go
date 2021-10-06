// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"

	"github.com/gin-gonic/gin"
)

func LoginRouteCollection(r *gin.Engine) *gin.Engine {
	r.GET("/login/teacher", controllers.TeacherLogin)
	r.GET("/login/student", controllers.StudentLogin)

	return r
}
