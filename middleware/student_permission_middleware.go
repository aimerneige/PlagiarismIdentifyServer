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

func StudentPermissionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get auth student from context
		authStudent, exist := c.Get("authStudent")
		if !exist {
			response.BadRequest(c, nil, "AuthStudent not found from context.")
			c.Abort()
			return
		}

		// get id from route
		id := c.Param("id")
		if id == "" {
			response.BadRequest(c, nil, "Student ID Required.")
			c.Abort()
			return
		}

		// access database with route id
		db := database.GetDB()
		var student models.Student
		db.First(&student, id)
		if student.ID == 0 {
			response.BadRequest(c, nil, "Student Not found")
			c.Abort()
			return
		}

		// if auth student has no permission
		if authStudent.(models.Student).ID != student.ID {
			response.BadRequest(c, nil, "No Permission")
			c.Abort()
			return
		}

		c.Next()
	}
}
