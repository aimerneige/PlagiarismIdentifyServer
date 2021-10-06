// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import (
	"time"

	"gorm.io/gorm"
)

type StudentHomework struct {
	gorm.Model
	StudentID    uint
	HomeworkID   uint
	FileItemIDs  []uint
	Upload       bool
	UploadTime   time.Time
	IsPlagiarism bool
}

type StudentHomeworkDto struct {
	ID           uint      `json:"id"`
	StudentID    uint      `json:"studentId"`
	HomeworkID   uint      `json:"homeworkId"`
	FileItemIDs  []uint    `json:"fileItemIds"`
	Upload       bool      `json:"upload"`
	UploadTime   time.Time `json:"uploadTime"`
	IsPlagiarism bool      `json:"isPlagiarism"`
}

func (s StudentHomework) ToDto() (dto StudentHomeworkDto) {
	dto.ID = s.ID
	dto.StudentID = s.StudentID
	dto.HomeworkID = s.HomeworkID
	dto.FileItemIDs = s.FileItemIDs // TODO test for bug
	dto.Upload = s.Upload
	dto.UploadTime = s.UploadTime
	dto.IsPlagiarism = s.IsPlagiarism

	return
}
