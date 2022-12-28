package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/absensi_restapi"))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Autoreply{})
	database.AutoMigrate(&Chat{})

	DB = database
}
