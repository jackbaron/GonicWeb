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
	c.String(200, "OK")
}

func BasicAuthenticateAdmin(c *gin.Context) {
	user := c.MustGet(gin.AuthUserKey).(string)
	if _, ok := secrets[user]; ok {
		sess := helpers.Instance(c.Request)
		if sess.Values["user"] != "" {
			c.Redirect(http.StatusMovedPermanently, "/admin/auth/login")
		} else {
			log.Println("admin")
		}
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
	session := helpers.Instance(c.Request)
	UserName := c.PostForm("UserName")
	PassWord := c.PostForm("PassWord")
	repo := dataservice.NewUserRepo()
	user := repo.GetUser(UserName)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(PassWord))
		if err == nil {
			session.Values["id"] = user.ID
			err = session.Save(c.Request, c.Writer)
			log.Println(session.Values["user"])
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate session token"})
			} else {
				log.Println("ok")
				// c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user"})
				c.Redirect(http.StatusMovedPermanently, "/admin")
			}
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}

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
		c.JSON(http.StatusInternalServerError, "Cannot insert user")
	}

	c.JSON(http.StatusOK, gin.H{"Flag": true, "Message": "Insert user successfully"})
}
