// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import "gorm.io/gorm"

type TaskFile struct {
	gorm.Model
	Name           string
	Path           string
	HomeworkTaskID uint
}

type TaskFileDto struct {
	ID             uint   `json:"id"`
	Name           string `json:"name"`
	Path           string `json:"path"`
	HomeworkTaskID uint   `json:"homeworkTaskId"`
}

func (f TaskFile) ToDto() (dto TaskFileDto) {
	dto.ID = f.ID
	dto.Name = f.Name
	dto.Path = f.Path
	dto.HomeworkTaskID = f.HomeworkTaskID

	return
}
