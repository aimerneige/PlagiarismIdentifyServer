// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"plagiarism-identify-server/bean"
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"

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

func TeacherInfoUpdate(c *gin.Context) {

}

func TeacherProfileUpdate(c *gin.Context) {

}
