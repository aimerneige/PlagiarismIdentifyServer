// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package middleware

import (
	"plagiarism-identify-server/bean"
	"plagiarism-identify-server/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegisterMiddleware Middleware for register
func RegisterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get and check user form data
		account := c.PostForm("account")
		if len(account) < 3 || len(account) > 24 {
			response.BadRequest(c, account, "Account length must between 3 and 24.")
			c.Abort()
			return
		}

		password := c.PostForm("password")
		if len(password) < 8 || len(password) > 16 {
			response.BadRequest(c, gin.H{
				"password": password,
				"length":   len(password),
			}, "Password length must between 8 and 16.")
			c.Abort()
			return
		}
		if isWeakPassword(password) {
			response.BadRequest(c, password, "Weak Password.")
			c.Abort()
			return
		}
		if password == account {
			response.BadRequest(c, nil, "Password can not be same with account.")
			c.Abort()
			return
		}
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			response.InternalServerError(c, err, "Crypt Unsuccessful.")
			c.Abort()
			return
		}

		// create a user
		registerUser := bean.RegisterUser{
			Account:  account,
			Password: string(hashPassword),
		}

		// write user into context
		c.Set("registerUser", registerUser)
		c.Next()
	}
}

// isWeakPassword simply check if password is weak password
func isWeakPassword(password string) bool {
	if password == "00000000" {
		return true
	}
	if password == "000000000" {
		return true
	}
	if password == "11111111" {
		return true
	}
	if password == "111111111" {
		return true
	}
	if password == "66666666" {
		return true
	}
	if password == "666666666" {
		return true
	}
	if password == "88888888" {
		return true
	}
	if password == "888888888" {
		return true
	}
	if password == "11223344" {
		return true
	}
	if password == "12345678" {
		return true
	}
	if password == "123456789" {
		return true
	}
	if password == "password" {
		return true
	}
	return false
}
