package controllers

import (
	"restful-template/database"
	"restful-template/models"
	"restful-template/response"
	"restful-template/token"
	"restful-template/utils"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	db := database.GetDB()

	name := c.PostForm("name")
	if len(name) < 3 || len(name) > 12 {
		response.BadRequest(c, nil, "User name length must between 3 and 12.")
		return
	}
	email := c.PostForm("email")
	if !utils.VerifyEmailFormat(email) {
		response.BadRequest(c, nil, "Invalid email.")
		return
	}
	phone := c.PostForm("phone")
	if !utils.VerifyChinaPhoneNumberFormat(phone) {
		response.BadRequest(c, nil, "Invalid phone.")
		return
	}
	password := c.PostForm("password")
	if len(password) < 8 || len(password) > 16 {
		response.BadRequest(c, nil, "Password length must between 8 and 16.")
		return
	}

	var user models.User
	db.Where("name = ?", name).First(&user)
	if user.ID != 0 {
		response.UnprocessableEntity(c, nil, "User exist.")
		return
	}
	db.Where("email = ?", email).First(&user)
	if user.ID != 0 {
		response.UnprocessableEntity(c, nil, "Mail exist.")
		return
	}
	db.Where("phone = ?", phone).First(&user)
	if user.ID != 0 {
		response.UnprocessableEntity(c, nil, "Phone exist.")
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.InternalServerError(c, err, "Crypt unsuccessful")
		return
	}
	newUser := models.User{
		Name:     name,
		Email:    email,
		Phone:    phone,
		Password: string(hashPassword),
		Super:    false,
	}
	db.Create(&newUser)

	response.OK(c, gin.H{
		"id":   newUser.ID,
		"name": newUser.Name,
	}, "Register successful!")
}

func Login(c *gin.Context) {
	db := database.GetDB()

	name := c.PostForm("name")
	if name == "" {
		response.BadRequest(c, nil, "Name Required.")
		return
	}
	password := c.PostForm("password")
	if len(password) < 6 || len(password) > 16 {
		response.BadRequest(c, nil, "Invalid Password.")
		return
	}

	var user models.User
	db.Where("name = ?", name).First(&user)
	if user.ID == 0 {
		response.NotFound(c, nil, "User doesn't exist.")
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		response.BadRequest(c, nil, "Wrong Password.")
		return
	}

	token, err := token.ReleaseToken(user, 2*time.Hour)
	if err != nil {
		response.InternalServerError(c, err, "Fail to release token.")
		return
	}

	response.OK(c, gin.H{
		"token": token,
	}, "Login Successful!")

}
