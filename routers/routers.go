package routers

import (
	"log"

	"github.com/gin-contrib/multitemplate"
	"github.com/hoangnhat/project/controllers"
	"github.com/hoangnhat/project/middlewares"

	"github.com/gin-gonic/gin"
)

const (
	URLPathViewAdmin           = "public/views/admin/"
	URLPathViewAdminPartials   = URLPathViewAdmin + "Partials/"
	URLPathViewAdminBreadcrumb = URLPathViewAdmin + "Breadcrumb/"
)

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
		authorized.GET("/categories/:type", controllers.CategoriesIndex)
		authorized.GET("/categories/:type/create", controllers.CategoriesCreateGET)
	}

	//! Router 404
	r.NoRoute(func(c *gin.Context) {
		c.HTML(200, "NotFound", gin.H{"Title": "Page Not Found"})
	})
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
	//*! Login and NotFound
	r.AddFromFiles("NotFound", URLPathViewAdmin+"404.html",
		URLPathViewAdminPartials+"Header.html",
		URLPathViewAdminPartials+"Footer.html",
		URLPathViewAdminPartials+"SideBar.html",
		URLPathViewAdmin+"Index.html")
	r.AddFromFiles("Login", URLPathViewAdmin+"auth/Login.html")

	//*! Home
	r.AddFromFiles("AdminHome", URLPathViewAdmin+"Base.html",
		URLPathViewAdminPartials+"Header.html",
		URLPathViewAdminPartials+"Footer.html",
		URLPathViewAdminPartials+"SideBar.html",
		URLPathViewAdmin+"Index.html")

	//*! Category
	r.AddFromFiles("CategoriesIndex", URLPathViewAdmin+"Base.html",
		URLPathViewAdminPartials+"Header.html",
		URLPathViewAdminPartials+"Footer.html",
		URLPathViewAdminPartials+"SideBar.html",
		URLPathViewAdminBreadcrumb+"Index.html",
		URLPathViewAdmin+"Categories/Index.html")
	r.AddFromFiles("CategoriesCreate", URLPathViewAdmin+"Base.html",
		URLPathViewAdminPartials+"Header.html",
		URLPathViewAdminPartials+"Footer.html",
		URLPathViewAdminPartials+"SideBar.html",
		URLPathViewAdminBreadcrumb+"Create.html",
		URLPathViewAdmin+"Categories/Create.html")

	return r
}
