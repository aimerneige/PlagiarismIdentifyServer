// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

type Student struct {
	User
	Course []Course `gorm:"many2many:course_students"`
}

type StudentDto struct {
	UserDto
	CourseIDs []uint `json:"courseIds"`
}

func (s Student) ToDto(dto StudentDto) {
	dto.ID = s.ID
	dto.Account = s.Account
	dto.Password = s.Password
	dto.ProfileID = s.Profile.ID
	dto.Name = s.Name
	dto.Phone = s.Phone
	dto.Email = s.Email
	for _, course := range s.Course {
		dto.CourseIDs = append(dto.CourseIDs, course.ID)
	}

	return
}
