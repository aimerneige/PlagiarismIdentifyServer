// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func FileRouteCollection(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	file := v1.Group("/file")

	taskFileUpload := file.Group("/task")
	taskFileUpload.Use(middleware.TeacherAuthMiddleware())
	taskFileUpload.POST("", controllers.TaskFileUpload)

	homeworkFileUpload := file.Group("/homework")
	homeworkFileUpload.Use(middleware.StudentAuthMiddleware())
	homeworkFileUpload.POST("", controllers.HomeworkFileUpload)

	return r
}
