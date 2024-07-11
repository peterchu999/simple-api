package model

import (
	"fmt"
	"log"
	"os"

	mysql "gorm.io/driver/mysql"
	sqlite "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func connectSqliteDatabase() {

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

func connectSqlDatabase() {

	var (
		username = os.Getenv("DATABASE_USERNAME")
		password = os.Getenv("DATABASE_PASSWORD")
		host     = os.Getenv("DATABASE_HOST")
		dbName   = os.Getenv("DATABASE_NAME")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, dbName)

	database, conErr := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if conErr != nil {
		panic(fmt.Sprintf("Failed to connect to database! %s", conErr.Error()))
	}

	log.Println("connect mysql successful")

	migrateErr := database.AutoMigrate(&Book{})

	if migrateErr != nil {
		panic(fmt.Sprintf("Migration Failed %s", migrateErr.Error()))
	}

	log.Println("migrate mysql successful")
	DB = database
}

func ConnectDatabase() {
	dbType := os.Getenv("DATABASE")
	log.Println(dbType)
	switch dbType {
	case "sql", "mysql":
		connectSqlDatabase()
	default:
		connectSqliteDatabase()
	}
}
