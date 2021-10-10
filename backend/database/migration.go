package database

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Material struct {
	ID int `gorm:"primaryKey" json:"id"`
	FileName string `json:"file_name"`
	Time string `json:"time"`
	URL string `json:"url"`
	Tags string `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func Migrate(conn *gorm.DB) {
	conn.AutoMigrate(&Material{})
}