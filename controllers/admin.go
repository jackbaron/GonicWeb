package controllers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/hoangnhat/project/models"

	"github.com/gin-gonic/gin"
)

// *****************************************************************************
// Admin
// *****************************************************************************

var secrets = gin.H{
	"admin": gin.H{"email": "thnhat94@gmail.com", "phone": "0868401501"},
}

func BasicAuthenticateAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if _, ok := secrets[user]; ok {
		c.Redirect(http.StatusMovedPermanently, "admin/auth/login")
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

func AdminLoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/auth/login.html", gin.H{
		"title": "Login",
	})
}

//TODO: Create token and sql create
func AdminRegisterPost(c *gin.Context) {
	passWord, _ := bcrypt.GenerateFromPassword([]byte("nhat1194"), bcrypt.DefaultCost)
	user := models.User{Email: "thnhat94@gmail.com", FullName: "HoangNhat", AliasName: "JPMe", Password: passWord}
}
