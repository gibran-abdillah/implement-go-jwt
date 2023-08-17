package models

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	db_name := os.Getenv("DB_NAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")

	// formatting db url for connecting to mysql

	DB_URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		db_user,
		db_password,
		db_host,
		db_port,
		db_name,
	)

	connector := mysql.Open(DB_URL)

	db, err := gorm.Open(connector)
	if err != nil {
		panic(err)
	}

	// start migrating

	db.AutoMigrate(&User{})
	DB = db
}
