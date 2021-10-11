// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"plagiarism-identify-server/controllers"

	"github.com/gin-gonic/gin"
)

func IndexRouteCollection(r *gin.Engine) *gin.Engine {
	r.GET("/", controllers.Index)

	return r
}
