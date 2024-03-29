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
	for {
		db.Where("course_code = ?", cid).First(&courseTemp)
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

	dto := course.ToDto()
	db := database.GetDB()
	var studentSlice []models.Student
	db.Model(&course).Association("Students").Find(&studentSlice)
	for _, student := range studentSlice {
		dto.StudentIDs = append(dto.StudentIDs, student.ID)
	}
	var homeworkTaskSlice []models.HomeworkTask
	db.Model(&course).Association("HomeworkTasks").Find(&homeworkTaskSlice)
	for _, task := range homeworkTaskSlice {
		dto.HomeworkTaskIDs = append(dto.HomeworkTaskIDs, task.ID)
	}

	response.OK(c, dto, "Get Course Info Successful.")
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

	db := database.GetDB()
	var studentSlice []models.Student
	var studentIds []uint
	db.Model(&course).Association("Students").Find(&studentSlice)
	for _, student := range studentSlice {
		studentIds = append(studentIds, student.ID)
	}

	response.OK(c, gin.H{
		"id":         course.ID,
		"studentIds": studentIds,
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

	// get id from route
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "Course ID Required.")
		return
	}
	// database obj
	db := database.GetDB()
	// check if course exist
	var course models.Course
	db.First(&course, id)
	if course.ID == 0 {
		response.NotFound(c, nil, "Course Not Found.")
		return
	}

	// student id from port form
	studentId := c.PostForm("studentId")
	// check if user exist
	var student models.Student
	db.First(&student, studentId)
	if student.ID == 0 {
		response.NotFound(c, nil, "Student Not Found.")
		return
	}

	// check user permission
	if !authUser.(bean.AuthUser).IsTeacher && authUser.(bean.AuthUser).UserID != student.ID {
		response.Forbidden(c, nil, "No Permission.")
		return
	}
	if authUser.(bean.AuthUser).IsTeacher && authUser.(bean.AuthUser).UserID != course.TeacherID {
		response.Forbidden(c, nil, "No Permission")
		return
	}

	// update relation
	var studentSlice []models.Student
	db.Model(&course).Association("Students").Find(&studentSlice)
	for _, stu := range studentSlice {
		if stu.ID == student.ID {
			response.BadRequest(c, nil, "Student exist in course.")
			return
		}
	}
	studentSlice = append(studentSlice, student)
	course.Students = studentSlice
	if err := db.Save(&course).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	// get student ids for return
	var studentIds []uint
	for _, stu := range course.Students {
		studentIds = append(studentIds, stu.ID)
	}

	response.OK(c, gin.H{
		"id":         course.ID,
		"studentIds": studentIds,
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

	// get id from route
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "Course ID Required.")
		return
	}
	// database obj
	db := database.GetDB()
	// check if course exist
	var course models.Course
	db.First(&course, id)
	if course.ID == 0 {
		response.NotFound(c, nil, "Course Not Found.")
		return
	}

	// student id from port form
	studentId := c.PostForm("studentId")
	// check if user exist
	var student models.Student
	db.First(&student, studentId)
	if student.ID == 0 {
		response.NotFound(c, nil, "Student Not Found.")
		return
	}

	// check user permission
	if !authUser.(bean.AuthUser).IsTeacher && authUser.(bean.AuthUser).UserID != student.ID {
		response.Forbidden(c, nil, "No Permission.")
		return
	}
	if authUser.(bean.AuthUser).IsTeacher && authUser.(bean.AuthUser).UserID != course.TeacherID {
		response.Forbidden(c, nil, "No Permission")
		return
	}

	// check if user exist in course
	var studentSlice []models.Student
	db.Model(&course).Association("Students").Find(&studentSlice)
	studentExist := false // check if student exist
	for _, stu := range studentSlice {
		if stu.ID == student.ID {
			studentExist = true
			break
		}
	}
	if !studentExist {
		response.BadRequest(c, nil, "Student not exist in course.")
		return
	}

	// save to database
	db.Model(&course).Association("Students").Delete(student)

	response.OK(c, gin.H{
		"id":        course.ID,
		"studentId": student.ID,
	}, "Remove Student Successful.")
}

func CourseTaskGet(c *gin.Context) {
	// course
	course, hasError := getCourseWithId(c)
	if hasError {
		return
	}

	// get task info
	var taskSlice []models.HomeworkTask
	var taskIds []uint
	err := database.GetDB().Model(&course).Association("HomeworkTasks").Find(&taskSlice)
	if err != nil {
		response.InternalServerError(c, err, "Database Association Error.")
		return
	}
	for _, task := range taskSlice {
		taskIds = append(taskIds, task.ID)
	}

	response.OK(c, gin.H{
		"id":              course.ID,
		"homeworkTaskIds": taskIds,
	}, "Course Task Get Successful.")
}

func CourseGetCourseWithCourseCode(c *gin.Context) {
	// get course code from post form
	courseCode := c.Query("code")
	if courseCode == "" {
		response.BadRequest(c, nil, "CourseCode Required.")
		return
	}

	// access database
	db := database.GetDB()
	var course models.Course
	db.Where("course_code = ?", courseCode).First(&course)
	if course.ID == 0 {
		response.NotFound(c, nil, "Course Not Found.")
		return
	}

	response.OK(c, gin.H{
		"id": course.ID,
	}, "Get Course With CourseCode Successful.")
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
