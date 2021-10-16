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

	// teacher from database
	var teacher models.Teacher
	db := database.GetDB()
	db.First(&teacher, authTeacher.(models.Teacher).ID)
	if teacher.ID == 0 {
		response.BadRequest(c, nil, "Teacher Not Exist.")
		return
	}

	// HomeworkTask Id from Query
	homeworkTaskId := c.Query("taskId")
	if homeworkTaskId == "" {
		response.BadRequest(c, nil, "HomeworkTask Id Required")
		return
	}

	// task from database
	var task models.HomeworkTask
	db.First(&task, homeworkTaskId)
	if task.ID == 0 {
		response.BadRequest(c, nil, "Task Not Exist.")
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
		err := os.MkdirAll(destDirectory, 0755)
		if err != nil {
			response.InternalServerError(c, err, "Fail To MkdirAll.")
			return
		}
	}
	if err := c.SaveUploadedFile(header, destFile); err != nil {
		response.InternalServerError(c, err, "Fail to save file into disk.")
		return
	}

	path := "/file/" + filePath

	taskFile := models.TaskFile{
		Name:           name,
		Path:           path,
		HomeworkTaskID: task.ID,
	}

	if err := db.Create(&taskFile).Error; err != nil {
		response.InternalServerError(c, err, "Database Create Error,")
		return
	}

	response.Created(c, gin.H{
		"id": taskFile.ID,
	}, "TaskFile Create Successful.")
}

func TaskFileGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "File Id Required.")
		return
	}

	var file models.TaskFile
	database.GetDB().First(&file, id)
	if file.ID == 0 {
		response.NotFound(c, nil, "File Not Found.")
		return
	}

	response.OK(c, file.ToDto(), "File Info Get Successful.")
}

func HomeworkFileUpload(c *gin.Context) {
	// get auth teacher from context
	authStudent, exist := c.Get("authStudent")
	if !exist {
		response.BadRequest(c, nil, "AuthTeacher not found from context.")
		return
	}

	// student from database
	var student models.Student
	db := database.GetDB()
	db.First(&student, authStudent.(models.Student).ID)
	if student.ID == 0 {
		response.BadRequest(c, nil, "Student Not Exist.")
		return
	}

	// Homework Id from Query
	homeworkId := c.Query("homeworkId")
	if homeworkId == "" {
		response.BadRequest(c, nil, "Student Homework Id Required")
		return
	}

	// homework from database
	var homework models.StudentHomework
	db.First(&homework, homeworkId)
	if homework.ID == 0 {
		response.BadRequest(c, nil, "Homework Not Exist.")
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
		err := os.MkdirAll(destDirectory, 0755)
		if err != nil {
			response.InternalServerError(c, err, "Fail To MkdirAll.")
			return
		}
	}
	if err := c.SaveUploadedFile(header, destFile); err != nil {
		response.InternalServerError(c, err, "Fail to save file into disk.")
		return
	}

	path := "/file/" + filePath

	homeworkFile := models.HomeworkFile{
		Name:              name,
		Path:              path,
		StudentHomeworkID: homework.ID,
	}

	homework.Upload = true
	homework.UploadTime = time.Now().Unix()

	if err := db.Create(&homeworkFile).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error,")
		return
	}

	response.Created(c, gin.H{
		"id": homeworkFile.ID,
	}, "HomeworkFile Create Successful.")
}

func HomeworkFileGet(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "File Id Required.")
		return
	}

	var file models.HomeworkFile
	database.GetDB().First(&file, id)
	if file.ID == 0 {
		response.NotFound(c, nil, "File Not Found.")
		return
	}

	response.OK(c, file.ToDto(), "File Info Get Successful.")
}
