package controllers

import (
	"restful-template/database"
	"restful-template/models"
	"restful-template/response"

	"github.com/gin-gonic/gin"
)

func CreateArticle(c *gin.Context) {
	db := database.GetDB()

	user, exist := c.Get("user")
	if !exist {
		response.BadRequest(c, nil, "User not found from contex.")
		return
	}
	article, exist := c.Get("article")
	if !exist {
		response.BadRequest(c, nil, "Article not found from contex.")
		return
	}

	newArticle := models.Article{
		UserID:  user.(models.User).ID,
		Title:   article.(models.Article).Title,
		Content: article.(models.Article).Content,
		Likes:   0,
	}

	db.Create(&newArticle)

	response.OK(c, gin.H{
		"id":    newArticle.ID,
		"title": newArticle.Title,
	}, "Create article success.")
}

func DeleteArticle(c *gin.Context) {
	db := database.GetDB()

	user, exist := c.Get("user")
	if !exist {
		response.BadRequest(c, nil, "User not found from contex.")
		return
	}

	userID := user.(models.User).ID
	articleID := c.PostForm("id")

	if len(articleID) == 0 {
		response.BadRequest(c, nil, "BadRequest")
		return
	}

	var article models.Article
	db.First(&article, articleID)
	if article.UserID != userID {
		response.BadRequest(c, nil, "BadRequest")
		return
	}

	db.Delete(&article)
	response.OK(c, gin.H{
		"id":    article.ID,
		"title": article.Title,
	}, "Delete Successful!")
}

func UpdateArticle(c *gin.Context) {
	db := database.GetDB()

	user, exist := c.Get("user")
	if !exist {
		response.BadRequest(c, nil, "User not found from contex.")
		return
	}

	userID := user.(models.User).ID
	articleID := c.PostForm("id")

	if len(articleID) == 0 {
		response.BadRequest(c, nil, "BadRequest")
		return
	}

	var article models.Article
	db.First(&article, articleID)
	if article.UserID != userID {
		response.BadRequest(c, nil, "BadRequest")
		return
	}

	postArticle, exist := c.Get("article")
	if !exist {
		response.BadRequest(c, nil, "Article not found from contex.")
		return
	}

	article.Title = postArticle.(models.Article).Title
	article.Content = postArticle.(models.Article).Content

	db.Save(&article)
	response.OK(c, gin.H{
		"id":    article.ID,
		"title": article.Title,
	}, "Update article successful!")
}

func GetArticle(c *gin.Context) {
	db := database.GetDB()

	articleID := c.PostForm("id")

	if len(articleID) == 0 {
		response.BadRequest(c, nil, "BadRequest")
		return
	}

	var article models.Article
	db.First(&article, articleID)
	if article.ID == 0 {
		response.NotFound(c, gin.H{
			"id": articleID,
		}, "NotFound")
		return
	}

	response.OK(c, article.ToDto(), "Get article successful!")
}
