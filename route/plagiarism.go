// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func PlagiarismRouteCollection(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	plagiarism := v1.Group("/plagiarism")

	getPlagiarismInfo := plagiarism.Group(":id/")
	getPlagiarismInfo.Use(middleware.TeacherAuthMiddleware())
	getPlagiarismInfo.GET("", controllers.GetPlagiarismInfo)

	return r
}
