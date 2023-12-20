package router

import (
	"github.com/TPM-Project-Larces/back-end.git/handler"
	"github.com/gin-gonic/gin"
)

func attestationRoutes(router *gin.Engine, basePath string, pathResource string) {

	attestation := router.Group(basePath + pathResource)
	{
		attestation.POST("/upload_challenge", handler.UploadChallenge)
		attestation.POST("/upload_signature", handler.UploadSignature)
		attestation.POST("/upload_attestation_key", handler.UploadAttestationKey)
	}
}
