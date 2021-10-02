// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title      string
	CourseCode string
	TeacherID  uint
	Students   []User `gorm:"many2many:course_students"`
}

type ClassDto struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	CourseCode string `json:"courseCode"`
	TeacherID  uint   `json:"teacherId"`
	StudentIDs []uint `json:"studentIds"`
}

func (c Course) ToDto() (dto ClassDto) {
	dto.TeacherID = c.TeacherID
	dto.Title = c.Title
	dto.CourseCode = c.CourseCode

	for _, student := range c.Students {
		dto.StudentIDs = append(dto.StudentIDs, student.ID)
	}

	return
}

func (c Course) GetStudentCount() (count int) {
	return len(c.Students)
}
