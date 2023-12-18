package router

import (
	"github.com/TPM-Project-Larces/back-end.git/handler"
	"github.com/gin-gonic/gin"
)

func fileRoutes(router *gin.Engine, basePath string, pathResource string) {

	file := router.Group(basePath + pathResource)
	{
		file.GET("", handler.GetFiles)
		file.GET("/by_name", handler.GetFileByName)
		file.GET("/by_username", handler.GetFilesByUsername)
		file.POST("/upload_encrypted_file", handler.SavedFile)
		file.POST("/upload_file", handler.UploadFile)
		file.DELETE("", handler.DeleteFile)
	}
}
