// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package controllers

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"plagiarism-identify-server/database"
	"plagiarism-identify-server/models"
	"plagiarism-identify-server/response"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

const (
	textHomework    = "文本作业"
	programHomework = "程序作业"
	pictureHomework = "图形作业"

	langTypeJava = "java"
	langTypeCPP  = "c/c++"
	langTypePy   = "python3"
	langTypeTxt  = "text"
	langTypeDoc  = "doc"
)

func GetPlagiarismInfo(c *gin.Context) {

	id := c.Param("id")
	if id == "" {
		response.BadRequest(c, nil, "HomeworkTask Id Required.")
		return
	}

	db := database.GetDB()
	var task models.HomeworkTask
	db.First(&task, id)
	if task.ID == 0 {
		response.NotFound(c, nil, "HomeworkTask Not Found.")
		return
	}

	rootPath := viper.GetString("common.path")
	algorithmRootPath := viper.GetString("algorithm.path")
	algorithmBinary := viper.GetString("algorithm.binary")
	algorithmOutputDir := viper.GetString("algorithm.outputDir")
	algorithmOutputFile := viper.GetString("algorithm.outputFile")

	outputPath := filepath.Join(algorithmRootPath, algorithmOutputDir)
	if err := os.RemoveAll(outputPath); err != nil {
		response.InternalServerError(c, err, "Error When Delete File.")
		return
	}

	app := "java"
	argsJar := "-jar"
	argsJarObj := filepath.Join(algorithmRootPath, algorithmBinary)
	argsPath := filepath.Join(rootPath, "task", id)

	var argsHomeworkForm string
	switch homeworkType := task.Type; homeworkType {
	case models.DocumentType:
		argsHomeworkForm = textHomework
	case models.ImageType:
		argsHomeworkForm = pictureHomework
	case models.ProgramType:
		argsHomeworkForm = programHomework
	default:
		argsHomeworkForm = textHomework
	}

	var argsLangType string
	switch languageType := task.Language; languageType {
	case models.NONE:
		argsLangType = langTypeTxt
	case models.JAVA:
		argsLangType = langTypeJava
	case models.C_CPP:
		argsLangType = langTypeCPP
	case models.PYTHON:
		argsLangType = langTypePy
	default:
		argsLangType = langTypeDoc
	}

	cmd := exec.Command(
		app,              // java
		argsJar,          // -jar
		argsJarObj,       // file.jar
		argsPath,         // file path
		argsHomeworkForm, // homework type
		argsLangType,     // homework language
	)
	_, err := cmd.Output()
	if err != nil {
		response.InternalServerError(c, err, "Error When Running Java Code.")
		return
	}

	outputFile := filepath.Join(algorithmRootPath, algorithmOutputDir, algorithmOutputFile)
	content, err := ioutil.ReadFile(outputFile)
	if err != nil {
		response.InternalServerError(c, err, "Error When Reading File.")
		return
	}

	response.OK(c, content, "Running Plagiarism Algorithm Successful.")
}
