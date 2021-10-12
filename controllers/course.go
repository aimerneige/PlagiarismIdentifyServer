// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"plagiarism-identify-server/bean"
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"
	"plagiarism-identify-server/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CourseCreate(c *gin.Context) {
	// get auth teacher from context
	authTeacher, exist := c.Get("authTeacher")
	if !exist {
		response.BadRequest(c, nil, "AuthTeacher not found from context.")
		return
	}

	// get title from context
	title := c.PostForm("title")
	if len(strings.TrimSpace(title)) == 0 {
		response.BadRequest(c, nil, "Title required.")
		return
	}

	db := database.GetDB()

	var courseTemp models.Course
	var cid string = utils.GenerateCid(6)
	for true {
		db.Where("CourseCode = ?", cid).First(&courseTemp)
		if courseTemp.ID == 0 {
			break
		}
		cid = utils.GenerateCid(6)
	}

	course := models.Course{
		Title:      title,
		CourseCode: cid,
		TeacherID:  authTeacher.(models.Teacher).ID,
	}

	if err := db.Create(&course).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.Created(c, gin.H{
		"id": course.ID,
	}, "Course Create Successful.")
}

func CourseInfoGet(c *gin.Context) {
	// course
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	response.OK(c, course.ToDto(), "Get Course Info Successful.")
}

func CourseInfoUpdate(c *gin.Context) {
	// course
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	// get title from post form
	title := c.PostForm("title")
	if len(strings.TrimSpace(title)) == 0 {
		response.BadRequest(c, nil, "Title required.")
		return
	}
	if title == course.Title {
		response.BadRequest(c, nil, "New title are same with old title.")
		return
	}

	course.Title = title

	// save to database
	if err := database.GetDB().Save(&course).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":    course.ID,
		"title": course.Title,
	}, "Course Info Update Successful.")
}

func CourseDelete(c *gin.Context) {
	// course
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	if err := database.GetDB().Delete(&course).Error; err != nil {
		response.InternalServerError(c, err, "Database Delete Error.")
		return
	}

	response.OK(c, gin.H{
		"id": course.ID,
	}, "Course Delete Successful.")
}

func CourseStudentGet(c *gin.Context) {
	// course
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	response.OK(c, gin.H{
		"id":         course.ID,
		"studentIds": course.ToDto().StudentIDs,
	}, "Course Student Get Successful.")
}

func CourseStudentCreate(c *gin.Context) {
	// get auth user
	authUser, exist := c.Get("authUser")
	if !exist {
		response.BadRequest(c, nil, "AuthUser not found from context.")
		c.Abort()
		return
	}

	// student id from port form
	studentId := c.PostForm("studentId")

	// check database
	db := database.GetDB()
	var student models.Student
	db.First(&student, studentId)
	if student.ID == 0 {
		response.NotFound(c, nil, "Student Not Found.")
		return
	}

	// check user permission
	if !authUser.(bean.AuthUser).IsTeacher && authUser.(bean.AuthUser).UserID != student.ID {
		response.BadRequest(c, nil, "No Permission.")
		return
	}

	// get course with id
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	course.Students = append(course.Students, student)

	// save to database
	if err := database.GetDB().Save(course).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":         course.ID,
		"studentIds": course.ToDto().StudentIDs,
	}, "Add New Student Successful.")
}

func CourseStudentDelete(c *gin.Context) {
	// get auth user
	authUser, exist := c.Get("authUser")
	if !exist {
		response.BadRequest(c, nil, "AuthUser not found from context.")
		c.Abort()
		return
	}

	// student id from port form
	studentId := c.PostForm("studentId")

	// check database
	db := database.GetDB()
	var student models.Student
	db.First(&student, studentId)
	if student.ID == 0 {
		response.NotFound(c, nil, "Student Not Found.")
		return
	}

	// check user permission
	if !authUser.(bean.AuthUser).IsTeacher && authUser.(bean.AuthUser).UserID != student.ID {
		response.BadRequest(c, nil, "No Permission.")
		return
	}

	// get course with id
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	// remove user
	for index, existStudent := range course.Students {
		if existStudent.ID == student.ID {
			course.Students = append(course.Students[:index], course.Students[index+1:]...)
			break
		}
	}

	// save to database
	if err := database.GetDB().Save(course).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":         course.ID,
		"studentIds": course.ToDto().StudentIDs,
	}, "Remove Student Successful.")
}

func CourseTaskGet(c *gin.Context) {
	// course
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	response.OK(c, gin.H{
		"id":              course.ID,
		"homeworkTaskIds": course.ToDto().HomeworkTaskIDs,
	}, "Course Task Get Successful.")
}

func CourseTaskCreate(c *gin.Context) {

}

func CourseTaskUpdate(c *gin.Context) {

}

func CourseTaskDelete(c *gin.Context) {

}

func getCourseWithId(c *gin.Context) (models.Course, bool) {
	var course models.Course

	// get id from route
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "Course ID Required.")
		return course, true
	}

	// access database
	database.GetDB().First(&course, id)
	if course.ID == 0 {
		response.NotFound(c, nil, "Course Not Found.")
		return course, true
	}

	return course, false
}
