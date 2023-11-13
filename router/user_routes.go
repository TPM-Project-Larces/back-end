package router

import (
	"github.com/TPM-Project-Larces/back-end.git/handler"
	"github.com/gin-gonic/gin"
)

func userRoutes(router *gin.Engine, basePath string, pathResource string) {

	user := router.Group(basePath + pathResource)
	{
		user.GET("", handler.GetUsers)
		user.GET("/username", handler.GetUserByUsername)
		user.POST("", handler.CreateUser)
		user.PUT("", handler.UpdateUser)
		user.DELETE("", handler.DeleteUser)
	}
}
