package router

import (
	docs "github.com/TPM-Project-Larces/agent.git/docs"
	"github.com/TPM-Project-Larces/agent.git/handler"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	basePath := "/"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)
	{
		//Show Oppening
		v1.POST("/upload_file", handler.UploadFile)
		v1.POST("/upload_key", handler.UploadKey)
		v1.GET("/decryptFile", handler.DecryptFile)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
