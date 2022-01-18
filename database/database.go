package database

import (
	"log"
	"os"
	"github.com/neerajbg/golang-fiber-boilerplate/model"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect(){

	dsn := os.Getenv("dsn")
  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})


	if err != nil {
	  panic("Error in connection")
	}else{
		log.Println("Connection successful.")
	}

	db.AutoMigrate(new(model.Blog))
	DB = db
}