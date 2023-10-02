package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary upload file

// @Description upload a file to encrypt
// @Tags Server operations
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} file_uploaded
// @Router /upload_file [post]
func UploadFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// Salva o arquivo no servidor
	err = ctx.SaveUploadedFile(file, "./files_to_encrypt/"+file.Filename)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("Arquivo recebido e armazenado com sucesso.")

	ctx.JSON(http.StatusOK, gin.H{"message": "Arquivo recebido e armazenado com sucesso"})
}
