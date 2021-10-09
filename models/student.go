// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Account  string
	Password string
	Avatar   string
	Name     string
	Phone    string
	Email    string
	Course   []Course `gorm:"many2many:course_students"`
}

type StudentDto struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CourseIDs []uint `json:"courseIds"`
}

func (s Student) ToDto() (dto StudentDto) {
	dto.ID = s.ID
	dto.Account = s.Account
	dto.Avatar = s.Avatar
	dto.Name = s.Name
	dto.Phone = s.Phone
	dto.Email = s.Email
	for _, course := range s.Course {
		dto.CourseIDs = append(dto.CourseIDs, course.ID)
	}

	return
}
