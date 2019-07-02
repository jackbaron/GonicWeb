package models

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// User struct
type User struct {
	ID        uint   `gorm:"primary_key"`
	UserName  string `gorm:"type:varchar(100);notn null;unique"`
	Email     string `gorm:"type:varchar(100);not null;unique"`
	Mobile    string `gorm:"type:varchar(16);not null;unique"`
	FullName  string `gorm:"not null"`
	AliasName string
	Password  []byte
	Avatar    string
	Level     int64 `gorm:"not null;default:0"`
	Token     []byte
	Role      int8 `gorm:"not null;default:1"`
	Status    int8 `gorm:"not null;default:0"`
	LastLogin time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

var hmacSecret = []byte{97, 48, 97, 50, 97, 98, 105, 49, 99, 102, 83, 53, 57, 98, 52, 54, 97, 102, 99, 12, 12, 13, 56, 34, 23, 16, 78, 67, 54, 34, 32, 21}

func EncodeToken(email string) string {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"nbf":   time.Now(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hmacSecret)

	fmt.Println(tokenString, err)

	return tokenString
}
