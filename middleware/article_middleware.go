package middleware

import (
	"restful-template/models"
	"restful-template/response"

	"github.com/gin-gonic/gin"
)

func ArticleMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		title := c.PostForm("title")
		content := c.PostForm("content")

		if len(title) > 64 {
			response.UnprocessableEntity(c, nil, "title length must less than 64.")
			c.Abort()
			return
		}

		if len(content) < 16 {
			response.UnprocessableEntity(c, nil, "content length must greater than 16.")
			c.Abort()
			return
		}

		article := models.Article{
			Title:   title,
			Content: content,
		}

		c.Set("article", article)
		c.Next()
	}
}
