package database

import (
	"os"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
	_ "github.com/joho/godotenv/autoload"
)

var db *gorm.DB

func Connect() *gorm.DB{

	dsn := os.Getenv("dsn")
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
	  panic("Error in connection")
	}
	return db
}