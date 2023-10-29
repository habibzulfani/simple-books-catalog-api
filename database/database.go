package database

import (
	"Mini-project/models"
	"Mini-project/utils"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	var err error

	var dsn string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.GetConfig("DB_USERNAME"),
		utils.GetConfig("DB_PASSWORD"),
		utils.GetConfig("DB_HOST"),
		utils.GetConfig("DB_PORT"),
		utils.GetConfig("DB_NAME"),
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("error when creating a connection to the database: %s\n", err)
	}

	log.Println("connected to database")
}

func DBMigration() {
	err := DB.AutoMigrate(models.User{})

	if err != nil {
		log.Fatalf("Failed to perform database migration: %s\n", err)
	}
}
