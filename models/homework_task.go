// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import (
	"gorm.io/gorm"
)

type HomeworkTask struct {
	gorm.Model
	Title            string
	Detail           string
	Type             HomeworkType
	Language         ProgramLanguage
	Files            []TaskFile
	DeadLine         int64
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
	CreateAt           int64           `json:"createAt"`
	UpdateAt           int64           `json:"updateAt"`
	DeadLine           int64           `json:"deadLine"`
	CourseID           uint            `json:"courseId"`
	StudentHomeworkIDs []uint          `json:"studentHomeworkIds"`
}

func (h HomeworkTask) ToDto() (dto HomeworkTaskDto) {
	dto.ID = h.ID
	dto.Title = h.Title
	dto.Detail = h.Detail
	dto.Type = h.Type
	dto.Language = h.Language
	dto.CreateAt = h.CreatedAt.Unix()
	dto.UpdateAt = h.UpdatedAt.Unix()
	dto.DeadLine = h.DeadLine
	dto.CourseID = h.CourseID

	return
}
