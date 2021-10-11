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

	teacherId := teacher.Group(":id/")

	teacherIdAuth := teacherId.Group("")
	teacherIdAuth.Use(middleware.TeacherAuthMiddleware())

	teacherIdAuth.GET("", controllers.TeacherInfoGet)

	teacherIdAuthPermission := teacherIdAuth.Group("")
	teacherIdAuthPermission.Use(middleware.TeacherPermissionMiddleware())

	teacherIdAuthPermission.PUT("", controllers.TeacherInfoUpdate)
	teacherIdAuthPermission.DELETE("", controllers.TeacherDelete)

	teacherIdAuthPermissionAvatar := teacherIdAuthPermission.Group("/avatar")
	teacherIdAuthPermissionAvatar.POST("", controllers.TeacherAvatarUpdate)

	teacherIdAuthAvatar := teacherIdAuth.Group("/avatar")
	teacherIdAuthAvatar.GET("", controllers.TeacherAvatarGet)

	teacherIdAuthPermissionName := teacherIdAuthPermission.Group("/name")
	teacherIdAuthPermissionName.PUT("", controllers.TeacherNameUpdate)

	teacherIdAuthPermissionPhone := teacherIdAuthPermission.Group("/phone")
	teacherIdAuthPermissionPhone.PUT("", controllers.TeacherPhoneUpdate)

	teacherIdAuthPermissionEmail := teacherIdAuthPermission.Group("/email")
	teacherIdAuthPermissionEmail.PUT("", controllers.TeacherEmailUpdate)

	teacherIdAuthPermissionPassword := teacherIdAuthPermission.Group("/password")
	teacherIdAuthPermissionPassword.PUT("", controllers.TeacherUpdatePassword)

	return r
}
