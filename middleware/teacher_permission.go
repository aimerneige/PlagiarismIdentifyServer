// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package middleware

import (
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"

	"github.com/gin-gonic/gin"
)

func TeacherPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get auth teacher from context
		authTeacher, exist := c.Get("authTeacher")
		if !exist {
			response.BadRequest(c, nil, "AuthTeacher not found from context.")
			c.Abort()
			return
		}

		// get id from route
		id := c.Param("id")
		if id == "" {
			response.BadRequest(c, nil, "Teacher ID Required.")
			c.Abort()
			return
		}

		// access database with route id
		db := database.GetDB()
		var teacher models.Teacher
		db.First(&teacher, id)
		if teacher.ID == 0 {
			response.BadRequest(c, nil, "Teacher Not found")
			c.Abort()
			return
		}

		// if auth teacher has no permission
		if authTeacher.(models.Teacher).ID != teacher.ID {
			response.BadRequest(c, nil, "No Permission")
			c.Abort()
			return
		}

		c.Next()
	}
}
