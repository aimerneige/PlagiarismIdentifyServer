package token

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID    uint
	IsTeacher bool
	jwt.StandardClaims
}
