// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"
	"plagiarism-identify-server/middleware"

	"github.com/gin-gonic/gin"
)

func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r = IndexRouteCollection(r)
	r = StaticRouteCollection(r)
	r = LoginRouteCollection(r)
	r = TeacherRouteCollection(r)
	r = StudentRouteCollection(r)
	r = CourseRouteCollection(r)
	r = TaskRouteCollection(r)
	r = HomeworkRouteCollection(r)
	r = FileRouteCollection(r)
	r = PlagiarismRouteCollection(r)

	r.NoRoute(controllers.NotFound)

	return r
}
