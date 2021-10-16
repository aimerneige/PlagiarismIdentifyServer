// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"

	"github.com/gin-gonic/gin"
)

func HomeworkInfoGet(c *gin.Context) {
	// get id from route
	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "HomeworkId required.")
		return
	}

	db := database.GetDB()

	var homework models.StudentHomework
	db.First(&homework, id)
	if homework.ID == 0 {
		response.NotFound(c, nil, "Homework Not Found.")
		return
	}

	var fileSlice []models.HomeworkFile
	err := db.Model(&homework).Association("Files").Find(&fileSlice)
	if err != nil {
		response.InternalServerError(c, err, "Database Association Error.")
		return
	}

	dto := homework.ToDto()
	for _, file := range fileSlice {
		dto.FileIds = append(dto.FileIds, file.ID)
	}

	response.OK(c, dto, "Homework Info Get Successful.")
}
