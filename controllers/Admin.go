package controllers

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"

	"github.com/hoangnhat/project/dataservice"

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

//TODO add CSRF post form login admin
func AdminLoginPOST(c *gin.Context) {
	session := sessions.Default(c)
	UserName := c.PostForm("UserName")
	PassWord := c.PostForm("PassWord")
	
}

func AdminRegisterPost(c *gin.Context) {
	var user models.User
	user.UserName = "admin"
	user.Email = "thnhat94@gmail.com"
	user.Mobile = "0868401501"
	user.FullName = "TaHoangNhat"
	user.AliasName = "JPME"
	user.Password, _ = bcrypt.GenerateFromPassword([]byte("nhat1194"), bcrypt.DefaultCost)
	repo := dataservice.NewUserRepo()
	err := repo.RegisterUser(&user)

	if !err {
		c.JSON(http.StatusInternalServerError, "Cannot insert iser")
	}

	c.JSON(http.StatusOK, gin.H{"Flag": true, "Message": "Insert user successfully"})
}
