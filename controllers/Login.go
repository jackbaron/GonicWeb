package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/hoangnhat/project/dataservice"
	"github.com/hoangnhat/project/helpers"
	"golang.org/x/crypto/bcrypt"
)

// store will hold all session data
var Store *sessions.CookieStore

const (
	// Name of the session variable that tracks login attempts
	sessLoginAttempt = "login_attempt"
)

//TODO add CSRF post form login admin
func AdminLoginPOST(c *gin.Context) {
	session := helpers.Instance(c.Request)
	UserName := c.PostForm("UserName")
	PassWord := c.PostForm("PassWord")

	// Check user submit post deveice
	if session.Values[sessLoginAttempt] != nil && session.Values[sessLoginAttempt].(int) >= 5 {
		log.Println("Brute force login prevented")
		AdminLoginGET(c)
		return
	}

	repo := dataservice.NewUserRepo()
	user := repo.GetUser(UserName)
	if user != nil {
		err := bcrypt.CompareHashAndPassword(user.Password, []byte(PassWord))
		if err == nil {
			helpers.Empty(session)
			session.Values["id"] = user.ID
			session.Values["UserName"] = user.UserName
			session.Values["Name"] = user.FullName
			session.Save(c.Request, c.Writer)
			log.Println("Successfully authenticated user")
			// c.JSON(http.StatusOK, gin.H{"message": "Successfully authenticated user", "ID": session.Values["id"]})
			c.Redirect(http.StatusMovedPermanently, "/admin/")
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
	}
}

func AdminLoginGET(c *gin.Context) {
	c.HTML(http.StatusOK, "Login", gin.H{
		"title": "Login",
	})
}

// loginAttempt increments the number of login attempts in sessions variable
func loginAttempt(sess *sessions.Session) {
	// Log the attempt
	if sess.Values[sessLoginAttempt] == nil {
		sess.Values[sessLoginAttempt] = 1
	} else {
		sess.Values[sessLoginAttempt] = sess.Values[sessLoginAttempt].(int) + 1
	}
}
