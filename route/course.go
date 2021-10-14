// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func CourseRouteCollection(r *gin.Engine) *gin.Engine {

	v1 := r.Group("/api/v1")
	course := v1.Group("/course")

	courseCreate := course.Group("")
	courseCreate.Use(middleware.TeacherAuthMiddleware())
	courseCreate.POST("", controllers.CourseCreate)

	courseCourseCode := course.Group("")
	courseCreate.Use(middleware.UserAuthMiddleware())
	courseCourseCode.GET("", controllers.CourseGetCourseWithCourseCode)

	courseId := course.Group(":id/")

	courseIdUserAuth := courseId.Group("")
	courseIdUserAuth.Use(middleware.UserAuthMiddleware())
	courseIdUserAuth.GET("", controllers.CourseInfoGet)

	courseIdTeacherAuth := courseId.Group("")
	courseIdTeacherAuth.Use(middleware.TeacherAuthMiddleware())
	courseIdTeacherAuth.PUT("", controllers.CourseInfoUpdate)
	courseIdTeacherAuth.DELETE("", controllers.CourseDelete)

	courseIdStudent := courseId.Group("/student")

	courseIdStudentUserAuth := courseIdStudent.Group("")
	courseIdStudentUserAuth.Use(middleware.UserAuthMiddleware())
	courseIdStudentUserAuth.GET("", controllers.CourseStudentGet)
	courseIdStudentUserAuth.POST("", controllers.CourseStudentCreate)
	courseIdStudentUserAuth.DELETE("", controllers.CourseStudentDelete)

	courseIdTask := courseId.Group("/task")

	courseIdTaskUserAuth := courseIdTask.Group("")
	courseIdTaskUserAuth.Use(middleware.UserAuthMiddleware())
	courseIdTaskUserAuth.GET("", controllers.CourseTaskGet)

	courseIdTaskTeacherAuth := courseIdTask.Group("")
	courseIdTaskTeacherAuth.Use(middleware.TeacherAuthMiddleware())
	courseIdTaskTeacherAuth.POST("", controllers.CourseTaskCreate)

	courseIdTaskUserAuthTaskId := courseIdTaskUserAuth.Group(":taskid/")
	courseIdTaskUserAuthTaskId.GET("", controllers.CourseTaskInfoGet)

	courseIdTaskTeacherAuthTaskId := courseIdTaskTeacherAuth.Group(":taskid/")
	courseIdTaskTeacherAuthTaskId.PUT("", controllers.CourseTaskUpdate)
	courseIdTaskTeacherAuthTaskId.DELETE("", controllers.CourseTaskDelete)

	return r
}
