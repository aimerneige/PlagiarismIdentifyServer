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
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func TaskCreate(c *gin.Context) {
	// get auth teacher from context
	authTeacher, exist := c.Get("authTeacher")
	if !exist {
		response.BadRequest(c, nil, "AuthTeacher not found from context.")
		return
	}

	// title
	title := c.PostForm("title")
	if len(strings.TrimSpace(title)) == 0 {
		response.BadRequest(c, nil, "Title required.")
		return
	}
	// detail
	detail := c.PostForm("detail")
	if len(strings.TrimSpace(detail)) == 0 {
		response.BadRequest(c, nil, "Detail required.")
		return
	}
	// homework
	homeworkType := c.PostForm("type")
	homeworkTypeValue, err := strconv.Atoi(homeworkType)
	if err != nil {
		response.BadRequest(c, gin.H{
			"err":          err,
			"homeworkType": homeworkType,
		}, "Wrong Args.")
		return
	}
	if homeworkTypeValue < 0 || homeworkTypeValue > 2 {
		response.BadRequest(c, nil, "Wrong Type.")
		return
	}
	// language
	language := c.PostForm("language")
	languageValue, err := strconv.Atoi(language)
	if err != nil {
		response.BadRequest(c, gin.H{
			"err":      err,
			"language": language,
		}, "Wrong Args.")
		return
	}
	if languageValue < 0 || languageValue > 3 {
		response.BadRequest(c, nil, "Wrong Type.")
		return
	}
	// time
	unixTime := c.PostForm("deadLine")
	deadLine, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		response.BadRequest(c, err, "Wrong Unix Time Format")
		return
	}
	// courseId from query
	courseId := c.Query("courseId")
	if courseId == "" {
		response.BadRequest(c, nil, "CourseId Required.")
		return
	}

	db := database.GetDB()

	// course
	var course models.Course
	db.First(&course, courseId)
	if course.ID == 0 {
		response.NotFound(c, nil, "Course Not Found.")
		return
	}

	// check permission
	if course.TeacherID != authTeacher.(models.Teacher).ID {
		response.Forbidden(c, nil, "No Permission.")
		return
	}

	// related student homework
	var studentSlice []models.Student
	var studentHomeworks []models.StudentHomework
	db.Model(&course).Association("Students").Find(&studentSlice)
	for _, student := range studentSlice {
		studentHomeworks = append(studentHomeworks, models.StudentHomework{
			StudentID:    student.ID,
			Upload:       false,
			IsPlagiarism: false,
		})
	}

	// create task
	task := models.HomeworkTask{
		Title:            title,
		Detail:           detail,
		Type:             models.HomeworkType(homeworkTypeValue),
		Language:         models.ProgramLanguage(languageValue),
		DeadLine:         deadLine,
		CourseID:         course.ID,
		StudentHomeworks: studentHomeworks,
	}

	// save to databases
	if err := db.Create(&task).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.Created(c, gin.H{
		"id": task.ID,
	}, "HomeworkTask Create Successful.")
}

func TaskInfoGet(c *gin.Context) {
	task, hasError := getTaskWithId(c)
	if hasError {
		return
	}

	dto := task.ToDto()
	db := database.GetDB()
	var fileSlice []models.TaskFile
	db.Model(&task).Association("TaskFile").Find(&fileSlice)
	for _, file := range fileSlice {
		dto.FileIDs = append(dto.FileIDs, file.ID)
	}
	var homeworkSlice []models.StudentHomework
	db.Model(&task).Association("StudentHomeworks").Find(&homeworkSlice)
	for _, homework := range homeworkSlice {
		dto.StudentHomeworkIDs = append(dto.StudentHomeworkIDs, homework.ID)
	}

	response.OK(c, dto, "HomeworkTask Info Get Successful.")
}

func TaskUpdate(c *gin.Context) {
	// TODO
	response.Forbidden(c, nil, "Not allowed to update.")
}

func TaskDelete(c *gin.Context) {
	// TODO
	response.Forbidden(c, nil, "Not allowed to delete.")
}

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

	// get task with id
	task, hasError := getTaskWithId(c)
	if hasError {
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

	var taskFileSlice []models.TaskFile
	if err := db.Model(&task).Association("Files").Find(&taskFileSlice); err != nil {
		response.InternalServerError(c, err, "Database Association Error.")
		return
	}
	taskFileSlice = append(taskFileSlice, taskFile)

	if err := db.Save(&task).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error,")
		return
	}

	var taskFileIds []uint
	for _, tf := range taskFileSlice {
		taskFileIds = append(taskFileIds, tf.ID)
	}

	response.OK(c, gin.H{
		"id":          task.ID,
		"taskFileIds": taskFileIds,
	}, "TaskFile Create Successful.")
}

func getTaskWithId(c *gin.Context) (models.HomeworkTask, bool) {
	var task models.HomeworkTask

	// get id from route
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "HomeworkTask ID Required.")
		return task, true
	}

	// access database
	database.GetDB().First(&task, id)
	if task.ID == 0 {
		response.NotFound(c, nil, "HomeworkTask Not Found.")
		return task, true
	}

	return task, false
}
