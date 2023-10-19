package router

import (
	docs "github.com/TPM-Project-Larces/back-end.git/docs"
	"github.com/TPM-Project-Larces/back-end.git/handler"
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
		v1.POST("/decrypt_file", handler.DecryptFile)
		v1.POST("/saved_file", handler.SavedFile)
	}
	v2 := router.Group(basePath)
	{
		v2.POST("create_user/", handler.CreateUser)

	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
