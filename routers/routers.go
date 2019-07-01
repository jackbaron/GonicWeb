package routers

import (
	"github.com/hoangnhat/project/controllers"

	"github.com/gin-gonic/gin"
)

/*
* get router running
 */
func SetRouter() *gin.Engine {
	r := gin.Default()
	// set Up Html Global
	r.LoadHTMLGlob("public/views/**/**/*")
	r.Static("/css", "public/assets/css")
	//* Router Admin
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "default",
	}))

	authorized.GET("/", controllers.BasicAuthenticateAdmin)
	authorized.GET("/auth/login", controllers.AdminLoginGET)

	r.POST("/register", controllers.AdminRegisterPost)

	return r
}
