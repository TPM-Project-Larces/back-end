package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary upload key

// @Description uploads a public key
// @Tags Server operations
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} key_uploaded
// @Router /upload_key [post]
func UploadKey(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Salva a chave no servidor
	err = ctx.SaveUploadedFile(file, "./key/"+file.Filename)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("Chave recebida e armazenada com sucesso.")

	ctx.JSON(http.StatusOK, gin.H{"message": "Chave pública recebida e armazenada com sucesso"})
}
