// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package models

import (
	"strings"

	"gorm.io/gorm"
)

const (
	IMAGE_JPG = ".jpg"
	IMAGE_PNG = ".png"
	IMAGE_GIF = ".gif"
	IMAGE_BMP = ".bmp"

	DOCUMENT_TXT  = ".txt"
	DOCUMENT_DOC  = ".doc"
	DOCUMENT_DOCX = ".docx"
	DOCUMENT_PDF  = ".pdf"
	DOCUMENT_HTML = ".html"

	PROGRAM_JAVA   = ".java"
	PROGRAM_C      = ".c"
	PROGRAM_CPP    = ".cpp"
	PROGRAM_PYTHON = ".py"
)

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

func (f HomeworkFile) IsImage() (ret bool) {
	ret = false
	fileName := strings.ToLower(f.Name)
	if strings.HasSuffix(fileName, IMAGE_JPG) {
		ret = true
	}
	if strings.HasSuffix(fileName, IMAGE_PNG) {
		ret = true
	}
	if strings.HasSuffix(fileName, IMAGE_GIF) {
		ret = true
	}
	if strings.HasSuffix(fileName, IMAGE_BMP) {
		ret = true
	}
	return
}

func (f HomeworkFile) IsDocument() (ret bool) {
	ret = false
	fileName := strings.ToLower(f.Name)
	if strings.HasSuffix(fileName, DOCUMENT_TXT) {
		ret = true
	}
	if strings.HasSuffix(fileName, DOCUMENT_DOC) {
		ret = true
	}
	if strings.HasSuffix(fileName, DOCUMENT_DOCX) {
		ret = true
	}
	if strings.HasSuffix(fileName, DOCUMENT_PDF) {
		ret = true
	}
	if strings.HasSuffix(fileName, DOCUMENT_HTML) {
		ret = true
	}
	return
}

func (f HomeworkFile) IsProgram() (ret bool) {
	ret = false
	fileName := strings.ToLower(f.Name)
	if strings.HasSuffix(fileName, PROGRAM_JAVA) {
		ret = true
	}
	if strings.HasSuffix(fileName, PROGRAM_C) {
		ret = true
	}
	if strings.HasSuffix(fileName, PROGRAM_CPP) {
		ret = true
	}
	if strings.HasSuffix(fileName, PROGRAM_PYTHON) {
		ret = true
	}
	return
}
