package models

import (
	"github.com/jinzhu/gorm"
)

type Content struct {
	gorm.Model
	Name        string `gorm:"not null"`
	InTro       string
	Content     string
	URL         string `gorm:"not null"`
	Status      int    `gorm:"not null;default:1"`
	Image       string
	Title       string   `gorm:"type:varchar(255)"`
	Description string   `gorm:"type:varchar(255)"`
	Category    Category `gorm:"foreignkey:CategoryID"`
	CategoryID  uint
	User        User `gorm:"foreignkey:UserID"`
	UserID      uint
	Order       int8
	Type        string `gorm:"type:varchar(100); not null"`
}
