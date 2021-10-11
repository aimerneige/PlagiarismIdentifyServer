// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package utils

import (
	"os"
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
