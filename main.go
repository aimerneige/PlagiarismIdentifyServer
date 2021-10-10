// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package main

import (
	"plagiarism-identify-server/config"
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/route"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	database.InitSQL()

	r := gin.Default()

	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r = route.AllRouteCollection(r)

	port := viper.GetString("common.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
