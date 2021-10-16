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

	taskFile := file.Group("/task")

	taskFileUpload := taskFile.Group("")
	taskFileUpload.Use(middleware.TeacherAuthMiddleware())
	taskFileUpload.POST("", controllers.TaskFileUpload)

	taskFileId := taskFile.Group(":id/")

	taskFileGet := taskFileId.Group("")
	taskFileGet.Use(middleware.UserAuthMiddleware())
	taskFileGet.GET("", controllers.TaskFileGet)

	homeworkFile := file.Group("/homework")

	homeworkFileUpload := homeworkFile.Group("")
	homeworkFileUpload.Use(middleware.StudentAuthMiddleware())
	homeworkFileUpload.POST("", controllers.HomeworkFileUpload)

	homeworkFileId := homeworkFile.Group(":id/")

	homeworkFileGet := homeworkFileId.Group("")
	homeworkFileGet.Use(middleware.UserAuthMiddleware())
	homeworkFileGet.GET("", controllers.HomeworkFileGet)

	return r
}
