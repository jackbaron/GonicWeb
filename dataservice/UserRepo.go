package dataservice

import (
	"log"

	"github.com/hoangnhat/project/models"
	"github.com/jinzhu/gorm"
)

//UserRepo struct
type UserRepo struct {
	db *gorm.DB
}

//NewUserRepo initialization connect DB
func NewUserRepo() *UserRepo {
	return &UserRepo{db: GetDBConection()}
}

//RegisterUser function insert user into DB
func (repo *UserRepo) RegisterUser(obj *models.User) bool {
	tx := repo.db.Begin()
	obj.Token = []byte(models.EncodeToken(obj.Email))
	err := tx.Create(obj).Error
	if err != nil {
		log.Fatal(err)
		log.Println("Cannot insert user")
		return false
	}

	tx.Commit()
	return true
}