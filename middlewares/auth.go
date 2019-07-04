package middlewares

import (
	"fmt"
	"net/http"

	"github.com/hoangnhat/project/helpers"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := helpers.Instance(c.Request)
		user := session.Values["user"]
		fmt.Println(user)
		if user == nil {
			// You'd normally redirect to login page
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid session token"})
			c.Redirect(http.StatusMovedPermanently, "/admin/auth/login")
		} else {
			// Continue down the chain to handler etc
			c.Next()
		}
	}
}
