// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package utils

import (
	"os"
	"strings"
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

// GetFileSize Get file size without open file
func GetFileSize(f string) (ret int64) {
	fi, err := os.Stat(f)
	if err != nil {
		return -1
	}
	return fi.Size()
}

// CheckFileExist Check if file exist
func CheckFileExist(f string) bool {
	if _, err := os.Stat(f); err == nil {
		return true
	}
	return false
}

// CheckFileIsDir Check if file a directory
func CheckFileIsDir(f string) bool {
	fi, err := os.Stat(f)
	if err != nil {
		return false
	}
	if fi.IsDir() {
		return true
	}
	return false
}

// CheckDirExist Check if directory exist
func CheckDirExist(f string) bool {
	if _, err := os.Stat(f); os.IsNotExist(err) {
		return false
	}
	return true
}

// IsImage Check if file is image
func IsImage(name string) (ret bool) {
	ret = false
	fileName := strings.ToLower(name)
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

// IsDocument Check if file is document
func IsDocument(name string) (ret bool) {
	ret = false
	fileName := strings.ToLower(name)
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

// IsProgram Check if file is program
func IsProgram(name string) (ret bool) {
	ret = false
	fileName := strings.ToLower(name)
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
