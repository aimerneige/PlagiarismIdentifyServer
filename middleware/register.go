// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package middleware

import (
	"plagiarism-identify-server/bean"
	"plagiarism-identify-server/response"
	"plagiarism-identify-server/utils"

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
		if utils.IsWeakPassword(password) {
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
