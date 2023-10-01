package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

// @BasePath /

// @Summary upload file

// @Description upload a file to encrypt
// @Tags User operations
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

	// Cria uma pasta para armazenar os arquivos, se necessário
	if err := os.MkdirAll("./encrypted_files", os.ModePerm); err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	// Salva o arquivo no servidor
	err = ctx.SaveUploadedFile(file, "./encrypted_files/"+file.Filename)
	if err != nil {
		SendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	fmt.Println("Arquivo recebido e armazenado com sucesso.")

	ctx.JSON(http.StatusOK, gin.H{"message": "Arquivo recebido e armazenado com sucesso"})
}
