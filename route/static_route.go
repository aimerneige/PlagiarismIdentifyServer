// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StaticRouteCollection(r *gin.Engine) *gin.Engine {
	path := viper.GetString("common.path")
	r.Static("/file", path)

	return r
}
