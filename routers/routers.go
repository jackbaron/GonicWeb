package routers

import (
	"log"

	"github.com/hoangnhat/project/controllers"
	"github.com/hoangnhat/project/middlewares"

	"github.com/gin-gonic/gin"
)

/*
* get router running
 */
func SetRouter() bool {
	r := gin.Default()
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// set Up Html Global
	r.LoadHTMLGlob("public/views/**/**/*")
	r.Static("/css", "public/assets/css")
	// Init Sessions
	//* Router Admin

	routerLogin := r.Group("manager")
	routerLogin.Use(middlewares.CheckAuthExist())
	{
		routerLogin.GET("/auth/login", controllers.AdminLoginGET)
		routerLogin.POST("/auth/login", controllers.AdminLoginPOST)
	}

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "default",
	}))
	authorized.GET("", controllers.BasicAuthenticateAdmin)
	authorized.Use(middlewares.AuthRequired())
	{
		authorized.GET("/", controllers.IndexHome)
		authorized.GET("/blog", controllers.IndexHome)
	}
	// r.GET("/register", controllers.AdminRegisterPost) // router insert user
	err := r.Run(":3500")
	if err != nil {
		log.Fatal("Error listening and server", err)
		return false
	}
	return true
}
