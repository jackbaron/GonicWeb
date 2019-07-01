package models

import "time"

// User struct
type User struct {
	ID        uint   `gorm:"primary_key"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Mobile    string `gorm:"type:varchar(16);not null;unique"`
	FullName  string `gorm:"not null"`
	AliasName string
	Password  []byte
	Avatar    string
	Level     int64  `gorm:"not null;default:0"`
	Token     string `gorm:"type:text"`
	Role      int8   `gorm:"not null;default:1"`
	Status    int8   `gorm:"not null;default:0"`
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
