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

	studentId := student.Group(":id/")

	studentIdAuth := studentId.Group("")
	studentIdAuth.Use(middleware.StudentAuthMiddleware())

	studentIdAuth.GET("", controllers.StudentInfoGet)

	studentIdAuthPermission := studentIdAuth.Group("")
	studentIdAuthPermission.Use(middleware.StudentPermissionMiddleware())

	studentIdAuthPermission.PUT("", controllers.StudentInfoUpdate)
	studentIdAuthPermission.DELETE("", controllers.StudentDelete)

	studentIdAuthPermissionAvatar := studentIdAuthPermission.Group("/avatar")
	studentIdAuthPermissionAvatar.POST("", controllers.StudentAvatarUpdate)

	studentIdAuthAvatar := studentIdAuth.Group("/avatar")
	studentIdAuthAvatar.GET("", controllers.StudentAvatarGet)

	studentIdAuthPermissionName := studentIdAuthPermission.Group("/name")
	studentIdAuthPermissionName.PUT("", controllers.StudentNameUpdate)

	studentIdAuthPermissionPhone := studentIdAuthPermission.Group("/phone")
	studentIdAuthPermissionPhone.PUT("", controllers.StudentPhoneUpdate)

	studentIdAuthPermissionEmail := studentIdAuthPermission.Group("/email")
	studentIdAuthPermissionEmail.PUT("", controllers.StudentEmailUpdate)

	studentIdAuthPermissionPassword := studentIdAuthPermission.Group("/password")
	studentIdAuthPermissionPassword.PUT("", controllers.StudentPasswordUpdate)

	return r
}
