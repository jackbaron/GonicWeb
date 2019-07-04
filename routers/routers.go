package routers

import (
	"github.com/hoangnhat/project/controllers"
	"github.com/hoangnhat/project/middlewares"

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
	// Init Sessions
	//* Router Admin
	// authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	"admin": "default",
	// }))
	// authorized.GET("/", controllers.BasicAuthenticateAdmin)
	authorized := r.Group("/admin")

	authorized.Use(middlewares.AuthRequired())
	{
		authorized.GET("/", controllers.IndexHome)
		authorized.GET("/auth/login", controllers.AdminLoginGET)
		authorized.POST("/auth/login", controllers.AdminLoginPOST)
	}
	// r.GET("/register", controllers.AdminRegisterPost) // router insert user
	return r
}
