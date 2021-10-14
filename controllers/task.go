// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func CourseTaskCreate(c *gin.Context) {
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
		response.BadRequest(c, nil, "Wrong Args.")
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
		response.BadRequest(c, nil, "Wrong Args.")
		return
	}
	if languageValue < 0 || languageValue > 3 {
		response.BadRequest(c, nil, "Wrong Type.")
		return
	}
	// time
	unixTime := c.PostForm("deadLine")
	unixTimeValue, err := strconv.ParseInt(unixTime, 10, 64)
	if err != nil {
		response.BadRequest(c, err, "Wrong Unix Time Format")
		return
	}
	deadLine := time.Unix(unixTimeValue, 0)
	// course
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	db := database.GetDB()

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
	}, "Task Create Successful.")
}

func CourseTaskInfoGet(c *gin.Context) {
	// TODO

	taskId := c.Param("taskid")
	if taskId == "" {
		response.BadRequest(c, nil, "TaskId Required.")
		return
	}

	response.OK(c, nil, "Task Info Get Successful.")
}

func CourseTaskUpdate(c *gin.Context) {
	// TODO
}

func CourseTaskDelete(c *gin.Context) {
	// TODO
}
