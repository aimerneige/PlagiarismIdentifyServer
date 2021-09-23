package middleware

import (
	"log"
	"restful-template/database"
	"restful-template/models"
	"restful-template/response"
	"restful-template/token"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")

		if tokenString == "" {
			response.NonAuthoritativeInfo(c, nil, "NonAuthoritativeInfo.")
			c.Abort()
			return
		}

		if !strings.HasPrefix(tokenString, "Bearer") {
			response.NoContent(c, nil, "Token String Invalid.")
			c.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := token.ParseToken(tokenString)

		if err != nil {
			response.Unauthorized(c, err, "Unauthorized.")
			c.Abort()
			return
		}
		if !token.Valid {
			response.Unauthorized(c, token, "Unauthorized.")
			c.Abort()
			return
		}
		log.Println("Authorized Successful.")

		userID := claims.UserID
		db := database.GetDB()
		var user models.User
		db.First(&user, userID)

		if userID == 0 {
			response.Forbidden(c, nil, "Forbidden")
			c.Abort()
			return
		}

		c.Set("user", user)
		c.Next()
	}
}
