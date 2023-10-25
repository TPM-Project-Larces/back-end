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

	auth := router.Group(basePath)
	{
		auth.POST("login/", handler.Login)
	}

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
		v2.POST("user/", handler.CreateUser)
		v2.GET("users/", handler.GetUsers)
		v2.GET("user/by_username", handler.GetUserByUsername)
		v2.DELETE("user/", handler.DeleteUser)
		v2.PUT("user/", handler.UpdateUser)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
