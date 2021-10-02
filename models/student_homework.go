package models

import "gorm.io/gorm"

type StudentHomework struct {
	gorm.Model
	StudentID  uint
	HomeworkID uint
	FileItemID uint
}

type StudentHomeworkDto struct {
	ID         uint
	StudentID  uint
	HomeworkID uint
}
