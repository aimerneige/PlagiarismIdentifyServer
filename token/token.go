// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package token

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const (
	ISSUER  = "aimerneige.com"
	SUBJECT = "user token"
)

var jwtKey = []byte("plagiarism-identify-server-key")

// ReleaseToken generate jwt token
func ReleaseToken(userId uint, isTeacher bool, tokenExpireDuration time.Duration) (string, error) {
	currentTime := time.Now()
	expirationTime := currentTime.Add(tokenExpireDuration)

	claims := &Claims{
		UserID:    userId,
		IsTeacher: isTeacher,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  currentTime.Unix(),
			Issuer:    ISSUER,
			Subject:   SUBJECT,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}

// ParseToken parse jwt token
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (i interface{}, err error) {
		return jwtKey, nil
	})

	return token, claims, err
}
