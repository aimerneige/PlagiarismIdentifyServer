// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"
	"plagiarism-identify-server/token"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func TeacherLogin(c *gin.Context) {
	// get and check user form data
	account := c.PostForm("account")
	if len(account) < 3 || len(account) > 24 {
		response.BadRequest(c, nil, "Invalid Required.")
		return
	}
	password := c.PostForm("password")
	if len(password) < 8 || len(password) > 16 {
		response.BadRequest(c, nil, "Invalid Password.")
		return
	}

	// check database
	db := database.GetDB()
	var teacher models.Teacher
	db.Where("account = ?", account).First(&teacher)
	if teacher.ID == 0 {
		response.NotFound(c, nil, "User doesn't exist.")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(password)); err != nil {
		response.BadRequest(c, nil, "Wrong Password.")
		return
	}

	// try to release token
	releaseToken, err := token.ReleaseToken(teacher.ID, true, 2*time.Hour)
	if err != nil {
		response.InternalServerError(c, err, "Fail to release releaseToken.")
		return
	}

	response.OK(c, gin.H{
		"id":           teacher.ID,
		"account":      teacher.Account,
		"releaseToken": releaseToken,
	}, "Login Successful.")
}

func StudentLogin(c *gin.Context) {
	// get and check user form data
	account := c.PostForm("account")
	if len(account) < 3 || len(account) > 24 {
		response.BadRequest(c, nil, "Invalid Required.")
		return
	}
	password := c.PostForm("password")
	if len(password) < 8 || len(password) > 16 {
		response.BadRequest(c, nil, "Invalid Password.")
		return
	}

	// check database
	db := database.GetDB()
	var student models.Student
	db.Where("account = ?", account).First(&student)
	if student.ID == 0 {
		response.NotFound(c, nil, "User doesn't exist.")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(password)); err != nil {
		response.BadRequest(c, nil, "Wrong Password.")
		return
	}

	// try to release token
	releaseToken, err := token.ReleaseToken(student.ID, false, 2*time.Hour)
	if err != nil {
		response.InternalServerError(c, err, "Fail to release releaseToken.")
		return
	}

	response.OK(c, gin.H{
		"id":           student.ID,
		"account":      student.Account,
		"releaseToken": releaseToken,
	}, "Login Successful.")
}
