package router

import (
	docs "github.com/TPM-Project-Larces/back-end.git/docs"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Initialize() {
	basePath := "/"
	docs.SwaggerInfo.BasePath = basePath

	router := gin.Default()
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	encryptionRoutes(router, basePath, "encryption/")

	router.Run(":5000")
}
