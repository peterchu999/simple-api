package model

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSqliteDatabase() {
	database, conErr := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if conErr != nil {
		panic(fmt.Sprintf("Failed to connect to database! %s", conErr.Error()))
	}
	log.Println("connect successful")
	migrateErr := database.AutoMigrate(&Book{})

	if migrateErr != nil {
		panic(fmt.Sprintf("Migration Failed %s", migrateErr.Error()))
	}

	log.Println("migrate successful")
	DB = database
}
