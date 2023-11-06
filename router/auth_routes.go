package router

import (
	"github.com/TPM-Project-Larces/back-end.git/handler"
	"github.com/gin-gonic/gin"
)

func authRoutes(router *gin.Engine, basePath string, pathResource string) {

	auth := router.Group(basePath + pathResource)
	{
		auth.POST("/login", handler.Login)
	}
}
