// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Account  string
	Password string
	Avatar   string
	Name     string
	Phone    string
	Email    string
	Courses  []Course
}

type TeacherDto struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Avatar    string `json:"avatar"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CourseIDs []uint `json:"courseIds"`
}

func (t Teacher) ToDto() (dto TeacherDto) {
	dto.ID = t.ID
	dto.Account = t.Account
	dto.Avatar = t.Avatar
	dto.Name = t.Name
	dto.Phone = t.Phone
	dto.Email = t.Email
	for _, course := range t.Courses {
		dto.CourseIDs = append(dto.CourseIDs, course.ID)
	}

	return
}
