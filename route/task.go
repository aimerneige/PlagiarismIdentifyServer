// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func TaskRouteCollection(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	task := v1.Group("/task")

	taskCreate := task.Group("")
	taskCreate.Use(middleware.TeacherAuthMiddleware())
	taskCreate.POST("", controllers.TaskCreate)

	taskId := task.Group(":id/")

	taskGet := taskId.Group("")
	taskGet.Use(middleware.UserAuthMiddleware())
	taskGet.GET("", controllers.TaskInfoGet)

	taskIdTeacherAuth := taskId.Group("")
	taskIdTeacherAuth.Use(middleware.TeacherAuthMiddleware())

	taskIdTeacherAuth.PUT("", controllers.TaskUpdate)
	taskIdTeacherAuth.DELETE("", controllers.TaskDelete)

	return r
}
