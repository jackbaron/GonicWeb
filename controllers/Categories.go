package controllers

import (
	"net/http"
	"strings"

	"github.com/hoangnhat/project/dataservice"

	"github.com/gin-gonic/gin"
)

//CategoriesIndex load list item category
func CategoriesIndex(c *gin.Context) {
	typeMethod := c.Param("type")
	Title := strings.Title(strings.Replace(typeMethod, "-", " ", -1))

	repo := dataservice.NewCategoriesRepo()
	listCategories := repo.GetListCategories(typeMethod)
	c.HTML(http.StatusOK, "CategoriesIndex", gin.H{"Title": Title, "Type": typeMethod, "listCategories": listCategories})
}

//CategoriesCreateGET create new item category
//TODO: Upload file. Validate file upload. Base64 filename
func CategoriesCreateGET(c *gin.Context) {
	typeMethod := c.Param("type")
	Category := strings.Title(strings.Replace(typeMethod, "-", " ", -1))
	repo := dataservice.NewCategoriesRepo()
	listCategory := repo.GetListCategories(typeMethod)
	Title := "Create Category " + Category
	c.HTML(http.StatusOK, "CategoriesCreate", gin.H{"Title": Title, "typeMethod": typeMethod, "Category": Category, "listCategory": listCategory})
}
