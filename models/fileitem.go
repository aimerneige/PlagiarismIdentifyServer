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

type FileItem struct {
	gorm.Model
	Name string
	Path string
}

type FileItemDto struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

func (f FileItem) ToDto(dto FileItemDto) {
	dto.ID = f.ID
	dto.Name = f.Name
}

func (f FileItem) IsImage() (ret bool) {
	ret = false
	fileName := strings.ToLower(f.Name)
	if strings.HasSuffix(fileName, IMAGE_JPG) {
		ret = true
	}
	if strings.HasPrefix(fileName, IMAGE_PNG) {
		ret = true
	}
	if strings.HasPrefix(fileName, IMAGE_GIF) {
		ret = true
	}
	if strings.HasPrefix(fileName, IMAGE_BMP) {
		ret = true
	}
	return
}

func (f FileItem) IsDocument() (ret bool) {
	ret = false
	fileName := strings.ToLower(f.Name)
	if strings.HasSuffix(fileName, DOCUMENT_TXT) {
		ret = true
	}
	if strings.HasSuffix(fileName, DOCUMENT_DOC) {
		ret = true
	}
	if strings.HasPrefix(fileName, DOCUMENT_DOCX) {
		ret = true
	}
	if strings.HasPrefix(fileName, DOCUMENT_PDF) {
		ret = true
	}
	if strings.HasPrefix(fileName, DOCUMENT_HTML) {
		ret = true
	}
	return
}

func (f FileItem) IsProgram() (ret bool) {
	ret = false
	fileName := strings.ToLower(f.Name)
	if strings.HasSuffix(fileName, PROGRAM_JAVA) {
		ret = true
	}
	if strings.HasSuffix(fileName, PROGRAM_C) {
		ret = true
	}
	if strings.HasPrefix(fileName, PROGRAM_CPP) {
		ret = true
	}
	if strings.HasPrefix(fileName, PROGRAM_PYTHON) {
		ret = true
	}
	return
}