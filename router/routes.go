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
		encryption.POST("/upload_encrypted_file", handler.SavedFile)
	}

	v3 := router.Group(basePath)
	{
		v3.GET("files/", handler.GetFiles)
		v3.GET("files/by_name/", handler.GetFileByName)
		v3.GET("files/by_username/", handler.GetFilesByUsername)
		v3.DELETE("files/", handler.DeleteFile)
	}
}
