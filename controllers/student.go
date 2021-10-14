// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"plagiarism-identify-server/bean"
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"
	"plagiarism-identify-server/utils"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
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

func StudentInfoGet(c *gin.Context) {
	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	var courseSlice []models.Course
	err := database.GetDB().Model(&student).Association("Courses").Find(&courseSlice)
	if err != nil {
		return
	}
	dto := student.ToDto()
	for _, course := range courseSlice {
		dto.CourseIDs = append(dto.CourseIDs, course.ID)
	}

	response.OK(c, dto, "Get User Info Successful.")
}

func StudentInfoUpdate(c *gin.Context) {
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

	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	student.Name = name
	student.Phone = phone
	student.Email = email

	// save to database
	if err := database.GetDB().Save(student).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":    student.ID,
		"name":  student.Name,
		"phone": student.Phone,
		"email": student.Email,
	}, "User Info Update Successful.")
}

func StudentAvatarGet(c *gin.Context) {
	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	response.OK(c, gin.H{
		"id":     student.ID,
		"avatar": student.Avatar,
	}, "Get User Avatar Successful.")
}

func StudentAvatarUpdate(c *gin.Context) {
	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	// avatar file
	file, header, err := c.Request.FormFile("avatar")
	if err != nil {
		response.BadRequest(c, err, "Fail to read form file.")
		return
	}
	defer file.Close()

	// check if file are image or not
	buff := make([]byte, 512)
	if _, err = file.Read(buff); err != nil {
		response.InternalServerError(c, err, "Fail to read avatar file into buff.")
		return
	}
	if contentType := http.DetectContentType(buff); !strings.HasPrefix(contentType, "image") {
		response.BadRequest(c, contentType, "Not a image file.")
		return
	}
	fileExt := filepath.Ext(header.Filename)
	if fileExt != ".png" && fileExt != ".jpg" && fileExt != ".jpeg" {
		response.BadRequest(c, nil, "Not a jpeg or png file.")
		return
	}

	// save file to disks
	rootPath := viper.GetString("common.path")                          // /opt/software/file
	fileName := "avatar" + fileExt                                      // avatar.jpg
	avatarDirectory := filepath.Join("avatar/student", student.Account) // avatar/student/1907040101
	destDirectory := filepath.Join(rootPath, avatarDirectory)           // /opt/software/file/avatar/student/1907040101
	avatarFile := filepath.Join(avatarDirectory, fileName)              // avatar/student/1907040101/avatar.jpg
	destFile := filepath.Join(destDirectory, fileName)                  // /opt/software/file/avatar/student/1907040101/avatar.jpg
	if !utils.CheckDirExist(destDirectory) {
		os.Mkdir(destDirectory, 0755)
	}
	if err := c.SaveUploadedFile(header, destFile); err != nil {
		response.InternalServerError(c, err, "Fail to save file to disk.")
		return
	}

	// avatar download link
	secure := viper.GetString("common.secure")
	baseUrl := viper.GetString("common.baseUrl")
	routePath := "file/"
	port := viper.GetString("common.port")
	if port == "80" {
		port = ""
	} else {
		port = ":" + port
	}
	avatarLink := fmt.Sprintf("%s://%s%s/%s%s",
		secure,
		baseUrl,
		port,
		routePath,
		avatarFile,
	)
	// http://example.com/file/avatar/student/1907040101/avatar.jpg

	// save to database
	student.Avatar = avatarLink

	if err := database.GetDB().Save(&student).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":     student.ID,
		"avatar": student.Avatar,
	}, "Update User Avatar Successful.")
}

func StudentNameUpdate(c *gin.Context) {
	// Name
	name := c.PostForm("name")
	if len(strings.TrimSpace(name)) == 0 {
		response.BadRequest(c, nil, "Name require.")
		c.Abort()
		return
	}

	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	student.Name = name

	// save to database
	if err := database.GetDB().Save(student).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":   student.ID,
		"name": student.Name,
	}, "User Name Update Successful.")
}

func StudentPhoneUpdate(c *gin.Context) {
	// Phone
	phone := c.PostForm("phone")
	if !utils.VerifyChinaPhoneNumberFormat(phone) {
		response.BadRequest(c, nil, "Wrong phone number!")
		c.Abort()
		return
	}

	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	student.Phone = phone

	// save to database
	if err := database.GetDB().Save(student).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":    student.ID,
		"phone": student.Phone,
	}, "User Phone Update Successful.")
}

func StudentEmailUpdate(c *gin.Context) {
	// Email
	email := c.PostForm("email")
	if !utils.VerifyEmailFormat(email) {
		response.BadRequest(c, nil, "Wrong email format!")
		c.Abort()
		return
	}

	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	student.Email = email

	// save to database
	if err := database.GetDB().Save(student).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":    student.ID,
		"email": student.Email,
	}, "User Email Update Successful.")
}

func StudentDelete(c *gin.Context) {
	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	if err := database.GetDB().Delete(&student).Error; err != nil {
		response.InternalServerError(c, err, "Database Delete Error.")
		return
	}

	response.OK(c, gin.H{
		"id": student.ID,
	}, "Delete Successful.")
}

func StudentPasswordUpdate(c *gin.Context) {
	// student
	student, hasError := getStudentWithId(c)
	if hasError {
		return
	}

	old := c.PostForm("old")
	new := c.PostForm("new")
	if old == new {
		response.BadRequest(c, nil, "New Password can not be same with old one.")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(student.Password), []byte(old)); err != nil {
		response.BadRequest(c, nil, "Wrong Password.")
		return
	}

	if len(new) < 8 || len(new) > 16 {
		response.BadRequest(c, gin.H{
			"password": new,
			"length":   len(new),
		}, "Password length must between 8 and 16.")
		c.Abort()
		return
	}
	if utils.IsWeakPassword(new) {
		response.BadRequest(c, new, "Weak Password.")
		c.Abort()
		return
	}
	if new == student.Account {
		response.BadRequest(c, nil, "Password can not be same with account.")
		c.Abort()
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(new), bcrypt.DefaultCost)
	if err != nil {
		response.InternalServerError(c, err, "Crypt Unsuccessful.")
		c.Abort()
		return
	}

	student.Password = string(hashPassword)

	// save to database
	if err := database.GetDB().Save(student).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":      student.ID,
		"account": student.Account,
	}, "User Password Update Successful.")
}

func getStudentWithId(c *gin.Context) (models.Student, bool) {
	var student models.Student

	// get id from route
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "Student ID Required.")
		return student, true
	}

	// access database
	database.GetDB().First(&student, id)
	if student.ID == 0 {
		response.NotFound(c, nil, "Student Not Found.")
		return student, true
	}

	return student, false
}
