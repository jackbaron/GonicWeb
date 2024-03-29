package controllers

import (
	"log"
	"net/http"

	"github.com/hoangnhat/project/helpers"

	"github.com/hoangnhat/project/dataservice"
	"github.com/hoangnhat/project/models"
	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

// *****************************************************************************
// Admin
// *****************************************************************************

var secrets = gin.H{
	"admin": gin.H{"email": "thnhat94@gmail.com", "phone": "0868401501"},
}

func IndexHome(c *gin.Context) {
	c.HTML(http.StatusOK, "AdminHome", gin.H{"Title": "Home"})
}

func BasicAuthenticateAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if _, ok := secrets[user]; ok {
		sess := helpers.Instance(c.Request)
		if sess.Values["user"] != "" {
			c.Redirect(http.StatusMovedPermanently, "/manager/auth/login")
		} else {
			log.Println("admin")
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
	}
}

func AdminRegisterPost(c *gin.Context) {
	var user models.User
	user.UserName = "admin"
	user.Email = "admin@gmail.com"
	user.Mobile = ""
	user.FullName = "TaHoangNhat"
	user.AliasName = "JPME"
	user.Password, _ = bcrypt.GenerateFromPassword([]byte("nhat1194"), bcrypt.DefaultCost)
	repo := dataservice.NewUserRepo()
	err := repo.RegisterUser(&user)

	if !err {
		c.JSON(http.StatusInternalServerError, "Cannot insert user")
	}

	c.JSON(http.StatusOK, gin.H{"Flag": true, "Message": "Insert user successfully"})
}
