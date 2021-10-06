// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package middleware

import (
	"net/http"
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"
	"plagiarism-identify-server/token"
	"strings"

	"github.com/gin-gonic/gin"
)

// StudentAuthMiddleware Middleware for student auth
func StudentAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check auth header
		statusCode, msg, data, claims := headerCheck(c)
		if statusCode != http.StatusOK {
			response.StatusCode(c, statusCode, data, msg)
			c.Abort()
			return
		}

		// get info from claims
		userID := claims.UserID
		isTeacher := claims.IsTeacher

		// if user are not student, return
		if isTeacher {
			response.BadRequest(c, nil, "Not a Student.")
			c.Abort()
			return
		}

		// check database
		db := database.GetDB()
		var student models.Student
		db.First(&student, userID)
		// check if user exist
		if userID == 0 {
			response.Forbidden(c, nil, "No Such User.")
			c.Abort()
			return
		}

		// write auth user into context
		c.Set("authUser", student)
		c.Next()
	}
}

// TeacherAuthMiddleware Middleware for teacher auth
func TeacherAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// check auth header
		statusCode, msg, data, claims := headerCheck(c)
		if statusCode != http.StatusOK {
			response.StatusCode(c, statusCode, data, msg)
			c.Abort()
			return
		}

		// get info from claims
		userID := claims.UserID
		isTeacher := claims.IsTeacher

		// if user are not teacher, return
		if !isTeacher {
			response.BadRequest(c, nil, "Not a Teacher.")
			c.Abort()
			return
		}

		// check database
		db := database.GetDB()
		var teacher models.Teacher
		db.First(&teacher, userID)
		// check if user exist
		if userID == 0 {
			response.Forbidden(c, nil, "No Such User.")
			c.Abort()
			return
		}

		// write auth user into context
		c.Set("authUser", teacher)
		c.Next()
	}
}

// headerCheck check header auth info
func headerCheck(c *gin.Context) (statusCode int, msg string, data interface{}, claims *token.Claims) {
	statusCode = http.StatusOK
	msg = ""

	// get parseToken from header
	tokenString := c.GetHeader("Authorization")

	// check parseToken string
	if tokenString == "" {
		statusCode = http.StatusNonAuthoritativeInfo
		msg = "NonAuthoritativeInfo."
		return
	}

	if !strings.HasPrefix(tokenString, "Bearer") {
		statusCode = http.StatusBadRequest
		msg = "Token String Invalid."
		return
	}

	tokenString = tokenString[7:]
	parseToken, claims, err := token.ParseToken(tokenString)

	if err != nil {
		statusCode = http.StatusUnauthorized
		data = err
		msg = "Fail To Parse Token."
		return
	}
	if !parseToken.Valid {
		statusCode = http.StatusUnauthorized
		data = parseToken
		msg = "Token Invalid."
		return
	}

	return
}
