// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import (
	"time"

	"gorm.io/gorm"
)

type Homework struct {
	gorm.Model
	Title    string
	Detail   string
	Type     HomeworkType
	Language ProgramLanguage
	Files    []FileItem
	DeadLine time.Time
}

type HomeworkDto struct {
	ID       uint            `json:"id"`
	Title    string          `json:"title"`
	Detail   string          `json:"detail"`
	Type     HomeworkType    `json:"type"`
	Language ProgramLanguage `json:"language"`
	FileIDs  []uint          `json:"fileIds"`
	CreateAt time.Time       `json:"createAt"`
	UpdateAt time.Time       `json:"updateAt"`
	DeadLine time.Time       `json:"deadLine"`
}

func (h Homework) ToDto(dto HomeworkDto) {
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

	return
}
