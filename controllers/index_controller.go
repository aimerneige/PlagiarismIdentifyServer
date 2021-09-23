package controllers

import (
	"restful-template/response"
	"time"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	data := gin.H{
		"ip":   c.ClientIP(),
		"time": time.Now().Local().Format("2006-01-02 15:04"),
	}
	msg := "Hello World!"
	response.OK(c, data, msg)
}
