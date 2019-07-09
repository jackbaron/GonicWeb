package routers

import (
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/hoangnhat/project/controllers"
	"github.com/hoangnhat/project/middlewares"

	"github.com/gin-gonic/gin"
)

const URLPathViewAdmin = "public/views/admin/"

func SetRouter() bool {
	r := gin.Default()
	r.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())
	// set Up Html Global
	r.Static("/assets", "public/assets")
	r.HTMLRender = buildTemplate()
	// Init Sessions
	//* Router Admin

	routerLogin := r.Group("manager")
	routerLogin.Use(middlewares.CheckAuthExist())
	{
		routerLogin.GET("/auth/login", controllers.AdminLoginGET)
		routerLogin.POST("/auth/login", controllers.AdminLoginPOST)
	}

	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"admin": "admin",
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

func buildTemplate() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromFiles("Login", URLPathViewAdmin+"auth/Login.html")
	r.AddFromFiles("AdminHome", URLPathViewAdmin+"Base.html", URLPathViewAdmin+"Partials/Header.html", URLPathViewAdmin+"Partials/Footer.html", URLPathViewAdmin+"Partials/SideBar.html", URLPathViewAdmin+"Index.html")
	return r
}
