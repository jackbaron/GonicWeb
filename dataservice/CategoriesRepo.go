package dataservice

import (
	"github.com/hoangnhat/project/models"
	"github.com/jinzhu/gorm"
)

type CategoriesRepo struct {
	db *gorm.DB
}

//NewCategoriesRepo connect DB
func NewCategoriesRepo() *CategoriesRepo {
	return &CategoriesRepo{db: GetDBConection()}
}

//GetListCategories get list category
func (repo *CategoriesRepo) GetListCategories(typeMethod string) (categories *models.Category) {
	listCategories := &models.Category{}
	err := repo.db.Where("type = ?", typeMethod).Find(listCategories)
	if err.RecordNotFound() {
		return nil
	}
	return
}
