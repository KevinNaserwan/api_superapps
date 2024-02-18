package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/API_superapps"))
	if err != nil {
		print(err)
	}

	database.AutoMigrate(&Product{})

	DB = database
}
