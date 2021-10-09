// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Title         string
	CourseCode    string
	TeacherID     uint
	Students      []Student `gorm:"many2many:course_students"`
	HomeworkTasks []HomeworkTask
}

type ClassDto struct {
	ID              uint   `json:"id"`
	Title           string `json:"title"`
	CourseCode      string `json:"courseCode"`
	TeacherID       uint   `json:"teacherId"`
	StudentIDs      []uint `json:"studentIds"`
	HomeworkTaskIDs []uint `json:"homeworkTaskIds"`
}

func (c Course) ToDto() (dto ClassDto) {
	dto.ID = c.ID
	dto.Title = c.Title
	dto.CourseCode = c.CourseCode
	dto.TeacherID = c.TeacherID
	for _, student := range c.Students {
		dto.StudentIDs = append(dto.StudentIDs, student.ID)
	}
	for _, homeworkTask := range c.HomeworkTasks {
		dto.HomeworkTaskIDs = append(dto.HomeworkTaskIDs, homeworkTask.ID)
	}

	return
}
