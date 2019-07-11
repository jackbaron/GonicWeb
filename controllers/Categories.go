package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/hoangnhat/project/dataservice"
	"github.com/hoangnhat/project/models"

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
func CategoriesCreateGET(c *gin.Context) {
	typeMethod := c.Param("type")
	Category := strings.Title(strings.Replace(typeMethod, "-", " ", -1))
	repo := dataservice.NewCategoriesRepo()
	listCategory := repo.GetListCategories(typeMethod)
	Title := "Create Category " + Category
	c.HTML(http.StatusOK, "CategoriesCreate", gin.H{"Title": Title, "typeMethod": typeMethod, "Category": Category, "listCategory": listCategory})
}

//TODO: Upload file. Validate file upload. Base64 filename
//CategoriesCreatePOST insert item category
func CategoriesCreatePOST(c *gin.Context) {
	category := models.Category{}
	category.Name = c.PostForm("name")
	category.Intro = c.PostForm("intro")
	category.Content = c.PostForm("content")
	category.Title = c.PostForm("title")
	category.Description = c.PostForm("description")
	category.Type = c.PostForm("type")
	file, _ := c.FormFile("image")
	if file != nil {
		if _, err := os.Stat("public/upload/" + category.Type); os.IsNotExist(err) {
			fmt.Println("create folder")
			os.Mkdir("public/upload/"+category.Type, 0755)
		}
		c.SaveUploadedFile(file, "public/upload/"+category.Type)

		category.Image = file.Filename
	}
	c.JSON(http.StatusBadRequest, gin.H{"error": category})
}
