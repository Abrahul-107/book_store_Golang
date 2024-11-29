package config

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Connect() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	dsn := user + ":" + password + "@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	log.Println("Database connection established!")
}

func GetDB() *gorm.DB {
	return db
}
