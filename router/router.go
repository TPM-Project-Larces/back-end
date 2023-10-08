package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Initialize Routes
	initializeRoutes(router)

	// Rodar a nossa API
	router.Run(":3000") //listen and serve on 0.0.0.0:3000
}
