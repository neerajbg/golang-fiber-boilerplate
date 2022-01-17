package model

import (
	"time"
	"github.com/neerajbg/golang-fiber-boilerplate/database"
)

var db = database.Connect()

type Blog struct {
	// gorm.Model
	ID    uint   `gorm:"primaryKey"`
	Title string `gorm:"column:title"`
	Post  string `gorm:"column:post"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func init(){
	db.AutoMigrate(new(Blog))
}