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
	StudentID      uint
	HomeworkTaskID uint
	Files          []HomeworkFile
	Upload         bool
	UploadTime     time.Time
	IsPlagiarism   bool
}

type StudentHomeworkDto struct {
	ID             uint      `json:"id"`
	StudentID      uint      `json:"studentId"`
	HomeworkTaskID uint      `json:"homeworkTaskId"`
	FileIds        []uint    `json:"fileIds"`
	Upload         bool      `json:"upload"`
	UploadTime     time.Time `json:"uploadTime"`
	IsPlagiarism   bool      `json:"isPlagiarism"`
}

func (s StudentHomework) ToDto() (dto StudentHomeworkDto) {
	dto.ID = s.ID
	dto.StudentID = s.StudentID
	dto.HomeworkTaskID = s.HomeworkTaskID
	for _, file := range s.Files {
		dto.FileIds = append(dto.FileIds, file.ID)
	}
	dto.Upload = s.Upload
	dto.UploadTime = s.UploadTime
	dto.IsPlagiarism = s.IsPlagiarism

	return
}
