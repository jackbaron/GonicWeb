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
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// set Up Html Global
	r.LoadHTMLGlob("public/views/**/**/*")
	r.Static("/css", "public/assets/css")
	// Init Sessions
	//* Router Admin
	// authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
	// 	"admin": "default",
	// }))
	// authorized.GET("/", controllers.BasicAuthenticateAdmin)
	r.GET("manager/auth/login", controllers.AdminLoginGET)
	r.POST("manager/auth/login", controllers.AdminLoginPOST)
	authorized := r.Group("/admin")
	authorized.Use(middlewares.AuthRequired())
	{
		authorized.GET("/", controllers.IndexHome)
		authorized.GET("/blog", controllers.IndexHome)
	}
	// r.GET("/register", controllers.AdminRegisterPost) // router insert user
	return r
}
