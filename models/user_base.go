// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

type User struct {
	ID       uint `gorm:"primarykey"`
	Account  string
	Password string
	Profile  FileItem
	Name     string
	Phone    string
	Email    string
}

type UserDto struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	ProfileID uint   `json:"profileId"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
}
