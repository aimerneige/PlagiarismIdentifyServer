// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
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

	var courseSlice []models.Course
	err := database.GetDB().Model(&teacher).Association("Courses").Find(&courseSlice)
	if err != nil {
		response.InternalServerError(c, err, "Database Association Error.")
		return
	}
	dto := teacher.ToDto()
	for _, course := range courseSlice {
		dto.CourseIDs = append(dto.CourseIDs, course.ID)
	}

	response.OK(c, dto, "Get User Info Successful.")
}

func TeacherInfoUpdate(c *gin.Context) {
	// Name
	name := c.PostForm("name")
	if len(strings.TrimSpace(name)) == 0 {
		response.BadRequest(c, nil, "Name require.")
		return
	}
	// Phone
	phone := c.PostForm("phone")
	if !utils.VerifyChinaPhoneNumberFormat(phone) {
		response.BadRequest(c, nil, "Wrong phone number!")
		return
	}
	// Email
	email := c.PostForm("email")
	if !utils.VerifyEmailFormat(email) {
		response.BadRequest(c, nil, "Wrong email format!")
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

func TeacherAvatarGet(c *gin.Context) {
	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	response.OK(c, gin.H{
		"id":     teacher.ID,
		"avatar": teacher.Avatar,
	}, "Get User Avatar Successful.")
}

func TeacherAvatarUpdate(c *gin.Context) {
	// teacher
	teacher, hasError := getTeacherWithId(c)
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
	avatarDirectory := filepath.Join("avatar/teacher", teacher.Account) // avatar/teacher/1907040101
	destDirectory := filepath.Join(rootPath, avatarDirectory)           // /opt/software/file/avatar/teacher/1907040101
	avatarFile := filepath.Join(avatarDirectory, fileName)              // avatar/teacher/1907040101/avatar.jpg
	destFile := filepath.Join(destDirectory, fileName)                  // /opt/software/file/avatar/teacher/1907040101/avatar.jpg
	if !utils.CheckDirExist(destDirectory) {
		err := os.MkdirAll(destDirectory, 0755)
		if err != nil {
			response.InternalServerError(c, err, "Fail To MkdirAll.")
			return
		}
	}
	if err := c.SaveUploadedFile(header, destFile); err != nil {
		response.InternalServerError(c, err, "Fail to save file to disk.")
		return
	}

	// avatar download link
	avatarLink := "/file/" + avatarFile
	// /file/avatar/teacher/1907040101/avatar.jpg

	// save to database
	teacher.Avatar = avatarLink

	if err := database.GetDB().Save(&teacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":     teacher.ID,
		"avatar": teacher.Avatar,
	}, "Update User Avatar Successful.")
}

func TeacherNameUpdate(c *gin.Context) {
	// Name
	name := c.PostForm("name")
	if len(strings.TrimSpace(name)) == 0 {
		response.BadRequest(c, nil, "Name require.")
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
		"name": teacher.Name,
	}, "User Name Update Successful.")
}

func TeacherPhoneUpdate(c *gin.Context) {
	// Phone
	phone := c.PostForm("phone")
	if !utils.VerifyChinaPhoneNumberFormat(phone) {
		response.BadRequest(c, nil, "Wrong phone number!")
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

	response.OK(c, gin.H{
		"id": teacher.ID,
	}, "Delete Successful.")
}

func TeacherPasswordUpdate(c *gin.Context) {
	// teacher
	teacher, hasError := getTeacherWithId(c)
	if hasError {
		return
	}

	old := c.PostForm("old")
	new := c.PostForm("new")
	if old == new {
		response.BadRequest(c, nil, "New Password can not be same with old one.")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(teacher.Password), []byte(old)); err != nil {
		response.BadRequest(c, nil, "Wrong Password.")
		return
	}

	if len(new) < 8 || len(new) > 16 {
		response.BadRequest(c, gin.H{
			"password": new,
			"length":   len(new),
		}, "Password length must between 8 and 16.")
		return
	}
	if utils.IsWeakPassword(new) {
		response.BadRequest(c, new, "Weak Password.")
		return
	}
	if new == teacher.Account {
		response.BadRequest(c, nil, "Password can not be same with account.")
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(new), bcrypt.DefaultCost)
	if err != nil {
		response.InternalServerError(c, err, "Crypt Unsuccessful.")
		return
	}

	teacher.Password = string(hashPassword)

	// save to database
	if err := database.GetDB().Save(teacher).Error; err != nil {
		response.InternalServerError(c, err, "Database Save Error.")
		return
	}

	response.OK(c, gin.H{
		"id":      teacher.ID,
		"account": teacher.Account,
	}, "User Password Update Successful.")
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
