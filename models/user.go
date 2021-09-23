package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string
	Email    string
	Phone    string
	Password string
	Super    bool
}

type UserDto struct {
	ID       uint   `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Super    bool   `json:"super"`
}

func (u User) ToDto() (dto UserDto) {
	dto.ID = u.ID
	dto.Name = u.Name
	dto.Email = u.Email
	dto.Phone = u.Phone
	dto.Password = u.Password
	dto.Super = u.Super

	return
}
