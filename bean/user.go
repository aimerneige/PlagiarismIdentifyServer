// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package bean

type RegisterUser struct {
	Account  string
	Password string
}

type AuthUser struct {
	UserID    uint
	IsTeacher bool
}
