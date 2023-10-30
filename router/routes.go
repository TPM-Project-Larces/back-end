package router

import (
	"github.com/TPM-Project-Larces/back-end.git/handler"
	"github.com/gin-gonic/gin"
)

func encryptionRoutes(router *gin.Engine, basePath string, pathResource string) {

	encryption := router.Group(basePath + pathResource)
	{
		encryption.POST("/upload_file", handler.UploadFile)
		encryption.POST("/upload_key", handler.UploadKey)
		encryption.POST("/decrypt_file", handler.DecryptFile)
		encryption.POST("/saved_file", handler.SavedFile)
	}
}
