// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import (
	"time"

	"gorm.io/gorm"
)

type HomeworkTask struct {
	gorm.Model
	Title            string
	Detail           string
	Type             HomeworkType
	Language         ProgramLanguage
	Files            []TaskFile
	DeadLine         time.Time
	CourseID         uint
	StudentHomeworks []StudentHomework
}

type HomeworkTaskDto struct {
	ID                 uint            `json:"id"`
	Title              string          `json:"title"`
	Detail             string          `json:"detail"`
	Type               HomeworkType    `json:"type"`
	Language           ProgramLanguage `json:"language"`
	FileIDs            []uint          `json:"fileIds"`
	CreateAt           time.Time       `json:"createAt"`
	UpdateAt           time.Time       `json:"updateAt"`
	DeadLine           time.Time       `json:"deadLine"`
	CourseID           uint            `json:"courseId"`
	StudentHomeworkIDs []uint          `json:"studentHomeworkIds"`
}

func (h HomeworkTask) ToDto() (dto HomeworkTaskDto) {
	dto.ID = h.ID
	dto.Title = h.Title
	dto.Detail = h.Detail
	dto.Type = h.Type
	dto.Language = h.Language
	for _, file := range h.Files {
		dto.FileIDs = append(dto.FileIDs, file.ID)
	}
	dto.CreateAt = h.CreatedAt
	dto.UpdateAt = h.UpdatedAt
	dto.DeadLine = h.DeadLine
	dto.CourseID = h.CourseID
	for _, studentHomework := range h.StudentHomeworks {
		dto.StudentHomeworkIDs = append(dto.StudentHomeworkIDs, studentHomework.ID)
	}

	return
}
