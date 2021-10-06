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

func StudentRegister(c *gin.Context) {
	// get registerUser from context
	registerUser, exist := c.Get("registerUser")
	if !exist {
		response.BadRequest(c, nil, "RegisterUser not found from context.")
		return
	}

	// check if account exist
	db := database.GetDB()
	var student models.Student
	db.Where("account = ?", registerUser.(bean.RegisterUser).Account).First(&student)
	if student.ID != 0 {
		response.UnprocessableEntity(c, nil, "User exist.")
		return
	}

	// create a student and write to database
	newStudent := models.Student{
		Account:  registerUser.(bean.RegisterUser).Account,
		Password: registerUser.(bean.RegisterUser).Password,
	}

	if err := db.Create(&newStudent).Error; err != nil {
		response.InternalServerError(c, err, "Database Create Error.")
		return
	}

	response.Created(c, gin.H{
		"id":      newStudent.ID,
		"account": newStudent.Account,
	}, "Register successful.")
}
