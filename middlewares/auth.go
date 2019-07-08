package middlewares

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hoangnhat/project/helpers"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := helpers.Instance(c.Request)
		fmt.Println(sess)
		if sess.Values["id"] == nil {
			// You'd normally redirect to login page
			// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
			c.Redirect(http.StatusMovedPermanently, "/manager/auth/login")
		} else {
			log.Println("ok redirect")
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}

func CheckAuthExist() gin.HandlerFunc {
	return func(c *gin.Context) {
		sess := helpers.Instance(c.Request)
		if sess.Values["id"] != nil {
			// You'd normally redirect to login page
			// c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
			c.Redirect(http.StatusMovedPermanently, "/admin")
		}
	}
}
