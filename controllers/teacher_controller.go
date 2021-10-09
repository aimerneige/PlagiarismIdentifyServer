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

func TeacherRegister(c *gin.Context) {
	// get registerUser from context
	registerUser, exist := c.Get("registerUser")
	if !exist {
		response.BadRequest(c, nil, "RegisterUser not found from context.")
		return
	}

	// check if account exist
	db := database.GetDB()
	var teacher models.Teacher
	db.Where("account = ?", registerUser.(bean.RegisterUser).Account).First(&teacher)
	if teacher.ID != 0 {
		response.UnprocessableEntity(c, nil, "User exist.")
		return
	}

	// create a teacher and write to database
	newTeacher := models.Teacher{
		Account:  registerUser.(bean.RegisterUser).Account,
		Password: registerUser.(bean.RegisterUser).Password,
	}

	if err := db.Create(&newTeacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Create Error.")
		return
	}

	response.Created(c, gin.H{
		"id":      newTeacher.ID,
		"account": newTeacher.Account,
	}, "Register successful.")
}

func TeacherInfoGet(c *gin.Context) {
	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	response.OK(c, teacher.ToDto(), "Get User Info Successful.")
}

func TeacherInfoUpdate(c *gin.Context) {
	// Name
	name := c.PostForm("name")
	if len(strings.TrimSpace(name)) == 0 {
		response.BadRequest(c, nil, "Name require.")
		c.Abort()
		return
	}
	// Phone
	phone := c.PostForm("phone")
	if !utils.VerifyChinaPhoneNumberFormat(phone) {
		response.BadRequest(c, nil, "Wrong phone number!")
		c.Abort()
		return
	}
	// Email
	email := c.PostForm("email")
	if !utils.VerifyEmailFormat(email) {
		response.BadRequest(c, nil, "Wrong email format!")
		c.Abort()
		return
	}

	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	teacher.Name = name
	teacher.Phone = phone
	teacher.Email = email

	// save to database
	if err := database.GetDB().Save(teacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":    teacher.ID,
		"name":  teacher.Name,
		"phone": teacher.Phone,
		"email": teacher.Email,
	}, "User Info Update Successful.")
}

func TeacherProfileUpdate(c *gin.Context) {
	// TODO
}

func TeacherNameUpdate(c *gin.Context) {
	// Name
	name := c.PostForm("name")
	if len(strings.TrimSpace(name)) == 0 {
		response.BadRequest(c, nil, "Name require.")
		c.Abort()
		return
	}

	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	teacher.Name = name

	// save to database
	if err := database.GetDB().Save(teacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":   teacher.ID,
		"name": teacher.ID,
	}, "User Name Update Successful.")
}

func TeacherPhoneUpdate(c *gin.Context) {
	// Phone
	phone := c.PostForm("phone")
	if !utils.VerifyChinaPhoneNumberFormat(phone) {
		response.BadRequest(c, nil, "Wrong phone number!")
		c.Abort()
		return
	}

	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	teacher.Phone = phone

	// save to database
	if err := database.GetDB().Save(teacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":    teacher.ID,
		"phone": teacher.Phone,
	}, "User Phone Update Successful.")
}

func TeacherEmailUpdate(c *gin.Context) {
	// Email
	email := c.PostForm("email")
	if !utils.VerifyEmailFormat(email) {
		response.BadRequest(c, nil, "Wrong email format!")
		c.Abort()
		return
	}

	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	teacher.Email = email

	// save to database
	if err := database.GetDB().Save(teacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":    teacher.ID,
		"email": teacher.Email,
	}, "User Email Update Successful.")
}

func TeacherDelete(c *gin.Context) {
	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	if err := database.GetDB().Delete(&teacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Delete Error.")
		return
	}

	response.NoContent(c, nil, "Delete Successful")
}

func getTeacherWithId(c *gin.Context) (models.Teacher, bool) {
	var teacher models.Teacher

	// get id from route
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "Teacher ID Required.")
		return teacher, true
	}

	// access database
	database.GetDB().First(&teacher, id)
	if teacher.ID == 0 {
		response.NotFound(c, nil, "Teacher Not Found.")
		return teacher, true
	}

	return teacher, false
}
