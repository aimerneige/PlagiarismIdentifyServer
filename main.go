package main

import (
	"restful-template/config"
	"restful-template/database"
	"restful-template/route"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	database.InitSQL()

	r := gin.Default()

	r.MaxMultipartMemory = 2 << 20

	r = route.AllRouteCollection(r)

	port := viper.GetString("common.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	panic(r.Run())
}
