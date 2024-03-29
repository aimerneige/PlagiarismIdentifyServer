// Copyright (c) 2021 AimerNeige
// aimer.neige@aimerneige.com
// All rights reserved.

package database

import (
	"fmt"
	"log"
	"plagiarism-identify-server/models"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB global db object
var DB *gorm.DB

// InitSQL init mysql database
func InitSQL() *gorm.DB {
	// get mysql config
	username := viper.Get("mysql.username")
	password := viper.Get("mysql.password")
	host := viper.Get("mysql.host")
	port := viper.Get("mysql.port")
	database := viper.Get("mysql.database")
	charset := viper.Get("mysql.charset")
	// format mysql config
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	// log for debug
	log.Printf("Database connect args: %s\n", args)

	// try to link mysql
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		log.Print("fail connect mysql: ", err)
	}

	// create table
	db.AutoMigrate(&models.Teacher{})
	db.AutoMigrate(&models.Student{})
	db.AutoMigrate(&models.Course{})
	db.AutoMigrate(&models.TaskFile{})
	db.AutoMigrate(&models.HomeworkFile{})
	db.AutoMigrate(&models.HomeworkTask{})
	db.AutoMigrate(&models.StudentHomework{})

	DB = db
	return db
}

// GetDB get global db
func GetDB() *gorm.DB {
	return DB
}
