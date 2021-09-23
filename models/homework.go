package models

import (
	"time"

	"gorm.io/gorm"
)

type Homework struct {
	gorm.Model
	Title    string
	Detail   string
	DeadLine time.Time
}
