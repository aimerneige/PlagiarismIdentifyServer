// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package token

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	UserID    uint
	IsTeacher bool
	jwt.StandardClaims
}
