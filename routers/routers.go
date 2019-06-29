package routers

import (
	"github.com/gin-gonic/gin"
)

// Set up Router run
func setUpRouter() {
	router := gin.Default()
	// Listen and Server in 0.0.0.0:8080
	router.Run("")
}
