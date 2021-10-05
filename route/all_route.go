// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import "github.com/gin-gonic/gin"

func AllRouteCollection(r *gin.Engine) *gin.Engine {
	r = IndexRouteCollection(r)
	r = UserRouteCollection(r)

	return r
}
