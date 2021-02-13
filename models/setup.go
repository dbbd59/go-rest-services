package models

import (
	"goRestServices/setting"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := setting.DbSetting.ConnUrl
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&Book{})

	DB = db
}
