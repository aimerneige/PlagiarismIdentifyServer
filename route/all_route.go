// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import "github.com/gin-gonic/gin"

func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r = IndexRouteCollection(r)
	r = LoginRouteCollection(r)
	r = TeacherRouteCollection(r)
	r = StudentRouteCollection(r)

	return r
}
