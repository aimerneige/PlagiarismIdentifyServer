// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package route

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func StaticRouteCollection(r *gin.Engine) *gin.Engine {
	path := viper.GetString("common.path")
	rootPath := filepath.Base(path)
	r.Static("/file", rootPath)

	return r
}
