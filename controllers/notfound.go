// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"plagiarism-identify-server/response"

	"github.com/gin-gonic/gin"
)

func NotFound(c *gin.Context) {
	response.NotFound(c, nil, "404 page not found")
}
