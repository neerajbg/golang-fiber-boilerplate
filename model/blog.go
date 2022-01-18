package model

import (
	"time"
)

type Blog struct {
	// gorm.Model
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title" gorm:"not null;column:title;size:256;"`
	Post  string `json:"post" gorm:"column:post;type:text"`
	CreatedAt time.Time
	UpdatedAt time.Time
}