package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusCode(c *gin.Context, statusCode int, data interface{}, msg string) {
	c.JSON(statusCode, gin.H{
		"code": statusCode,
		"data": data,
		"msg":  msg,
	})
}

// OK 200
func OK(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusOK, data, msg)
}

// Created 201
func Created(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusCreated, data, msg)
}

// NonAuthoritativeInfo 203
func NonAuthoritativeInfo(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusNonAuthoritativeInfo, data, msg)
}

// NoContent 204
func NoContent(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusNoContent, data, msg)
}

// BadRequest 400
func BadRequest(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusBadRequest, data, msg)
}

// Unauthorized 401
func Unauthorized(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusUnauthorized, data, msg)
}

// Forbidden 403
func Forbidden(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusForbidden, data, msg)
}

// NotFound 404
func NotFound(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusNotFound, data, msg)
}

// UnprocessableEntity 422
func UnprocessableEntity(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusUnprocessableEntity, data, msg)
}

// InternalServerError 500
func InternalServerError(c *gin.Context, data interface{}, msg string) {
	StatusCode(c, http.StatusInternalServerError, data, msg)
}
