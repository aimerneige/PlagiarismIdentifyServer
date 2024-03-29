// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import (
	"gorm.io/gorm"
)

type StudentHomework struct {
	gorm.Model
	StudentID      uint
	HomeworkTaskID uint
	Files          []HomeworkFile
	Upload         bool
	UploadTime     int64
	IsPlagiarism   bool
}

type StudentHomeworkDto struct {
	ID             uint   `json:"id"`
	StudentID      uint   `json:"studentId"`
	HomeworkTaskID uint   `json:"homeworkTaskId"`
	FileIds        []uint `json:"fileIds"`
	Upload         bool   `json:"upload"`
	UploadTime     int64  `json:"uploadTime"`
}

func (s StudentHomework) ToDto() (dto StudentHomeworkDto) {
	dto.ID = s.ID
	dto.StudentID = s.StudentID
	dto.HomeworkTaskID = s.HomeworkTaskID
	dto.Upload = s.Upload
	dto.UploadTime = s.UploadTime

	return
}
