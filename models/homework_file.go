// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import "gorm.io/gorm"

type HomeworkFile struct {
	gorm.Model
	Name              string
	Path              string
	StudentHomeworkID uint
}

type HomeworkFileDto struct {
	ID                uint   `json:"id"`
	Name              string `json:"name"`
	Path              string `json:"path"`
	StudentHomeworkID uint   `json:"studentHomeworkId"`
}

func (f HomeworkFile) ToDto() (dto HomeworkFileDto) {
	dto.ID = f.ID
	dto.Name = f.Name
	dto.Path = f.Path
	dto.StudentHomeworkID = f.StudentHomeworkID

	return
}
