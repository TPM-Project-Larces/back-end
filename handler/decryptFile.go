package handler

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary decrypt file

// @Description upload a file to decrypt
// @Tags Server operations
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {array} string "Folders available for upload"
// @Router /decrypt_file [post]
func DecryptFile(ctx *gin.Context) {
	// Obtenha o nome da pasta escolhida pelo usuário a partir dos parâmetros da consulta
	folder := ctx.DefaultQuery("folder", "")

	if folder == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "O nome da pasta é obrigatório"})
		return
	}

	// Verifica se a pasta é válida e existe no diretório de upload
	uploadFolder := "./encrypted_files"
	if !isValidFolder(folder, uploadFolder) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "A pasta escolhida não é válida"})
		return
	}

	// Recebe o arquivo enviado pelo cliente
	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Caminho completo do arquivo no sistema de arquivos
	filePath := filepath.Join(uploadFolder, folder, file.Filename)

	// Faz o upload do arquivo para o sistema de arquivos
	if err := ctx.SaveUploadedFile(file, filePath); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	url := "http://localhost:3000/decrypt_file/"

	sendFile(filePath, url)

	ctx.JSON(http.StatusOK, gin.H{"message": "Encrypted file successfully sent to the client"})
}

// Função para verificar se a pasta é válida
func isValidFolder(folder, uploadFolder string) bool {
	// Verifique se a pasta existe no diretório de upload
	folderPath := filepath.Join(uploadFolder, folder)
	_, err := os.Stat(folderPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
