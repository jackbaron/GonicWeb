package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name        string    `gorm:"not null"`
	URL         string    `gorm:"not null"`
	Parent      *Category `gorm:"foreignkey:ParentID"`
	ParentID    uint
	Intro       string
	Content     string
	Status      uint8 `gorm:"not null; default(1)"`
	User        User  `gorm:"foreignkey:UserID"`
	UserID      uint
	Image       string
	Order       int8
	Title       string `gorm:"varchar(255)"`
	Description string `gorm:"varchar(255)"`
	Type        string `gorm:"type:varchar(100); not null"`
}
