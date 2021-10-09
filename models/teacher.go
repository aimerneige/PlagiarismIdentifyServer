// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Account  string
	Password string
	Profile  string
	Name     string
	Phone    string
	Email    string
	Courses  []Course
}

type TeacherDto struct {
	ID        uint   `json:"id"`
	Account   string `json:"account"`
	Password  string `json:"password"`
	Profile   string `json:"profile"`
	Name      string `json:"name"`
	Phone     string `json:"phone"`
	Email     string `json:"email"`
	CourseIDs []uint `json:"courseIds"`
}

func (t Teacher) ToDto() (dto TeacherDto) {
	dto.ID = t.ID
	dto.Account = t.Account
	dto.Password = t.Password
	// dto.Profile = "TODO DOWNLOAD LINK" // TODO
	dto.Name = t.Name
	dto.Phone = t.Phone
	dto.Email = t.Email
	for _, course := range t.Courses {
		dto.CourseIDs = append(dto.CourseIDs, course.ID)
	}

	return
}
