// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"os"
	"path/filepath"
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"
	"plagiarism-identify-server/utils"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TaskFileUpload(c *gin.Context) {
	// get auth teacher from context
	authTeacher, exist := c.Get("authTeacher")
	if !exist {
		response.BadRequest(c, nil, "AuthTeacher not found from context.")
		return
	}

	var teacher models.Teacher
	db := database.GetDB()
	db.First(&teacher, authTeacher.(models.Teacher).ID)
	if teacher.ID == 0 {
		response.BadRequest(c, nil, "Teacher Not Exist.")
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, err, "Fail to read form file.")
		return
	}
	defer file.Close()

	name := filepath.Base(header.Filename)
	fileExt := filepath.Ext(header.Filename)
	randomChars := utils.GenerateCid(16)
	rootPath := viper.GetString("common.path") // /opt/software/file
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "-" + randomChars + fileExt
	fileDirectory := filepath.Join("task", teacher.Account)
	destDirectory := filepath.Join(rootPath, fileDirectory)
	filePath := filepath.Join(fileDirectory, fileName)
	destFile := filepath.Join(destDirectory, fileName)
	if !utils.CheckDirExist(destDirectory) {
		os.Mkdir(destDirectory, 0755)
	}
	if err := c.SaveUploadedFile(header, destFile); err != nil {
		response.InternalServerError(c, err, "Fail to save file into disk.")
		return
	}

	path := "/file/" + filePath

	taskFile := models.TaskFile{
		Name: name,
		Path: path,
	}

	if err := db.Create(&taskFile).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error,")
		return
	}

	response.Created(c, gin.H{
		"id": taskFile.ID,
	}, "TaskFile Create Successful.")
}

func HomeworkFileUpload(c *gin.Context) {
	// get auth teacher from context
	authStudent, exist := c.Get("authStudent")
	if !exist {
		response.BadRequest(c, nil, "AuthTeacher not found from context.")
		return
	}

	var student models.Student
	db := database.GetDB()
	db.First(&student, authStudent.(models.Student).ID)
	if student.ID == 0 {
		response.BadRequest(c, nil, "Student Not Exist.")
		return
	}

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		response.BadRequest(c, err, "Fail to read form file.")
		return
	}
	defer file.Close()

	name := filepath.Base(header.Filename)
	fileExt := filepath.Ext(header.Filename)
	randomChars := utils.GenerateCid(16)
	rootPath := viper.GetString("common.path") // /opt/software/file
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "-" + randomChars + fileExt
	fileDirectory := filepath.Join("homework", student.Account)
	destDirectory := filepath.Join(rootPath, fileDirectory)
	filePath := filepath.Join(fileDirectory, fileName)
	destFile := filepath.Join(destDirectory, fileName)
	if !utils.CheckDirExist(destDirectory) {
		os.Mkdir(destDirectory, 0755)
	}
	if err := c.SaveUploadedFile(header, destFile); err != nil {
		response.InternalServerError(c, err, "Fail to save file into disk.")
		return
	}

	path := "/file/" + filePath

	homeworkFile := models.HomeworkFile{
		Name: name,
		Path: path,
	}

	if err := db.Create(&homeworkFile).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error,")
		return
	}

	response.Created(c, gin.H{
		"id": homeworkFile.ID,
	}, "HomeworkFile Create Successful.")
}