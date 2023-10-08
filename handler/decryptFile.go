package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

// @Summary Decrypt a file
// @Description Provide the filename to decrypt
// @Tags Server operations
// @Produce json
// @Param filename formData string true "Filename to decrypt"
// @Success 200 {object} string "Decrypted file"
// @Router /decrypt_file [post]
func DecryptFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	nameFile := ctx.PostForm("filename")

	uploadDir := "./encrypted_files"
	filePath := filepath.Join(uploadDir, nameFile)
	fmt.Println(filePath)
	_, err := os.Stat(filePath)

	if nameFile == "" || os.IsNotExist(err) {
		// O arquivo não existe, retorna uma mensagem de erro
		ctx.JSON(http.StatusNotFound, gin.H{"error": "not found file"})
		return
	} else if err != nil {
		// Outro erro ocorreu, retorna um erro interno do servidor
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Retorna uma mensagem de sucesso.
	ctx.JSON(http.StatusOK, gin.H{"message": "file found and sent to another API"})

	url := "http://localhost:5000/decrypt_file/"

	sendFile(filePath, url)

	ctx.JSON(http.StatusOK, gin.H{"message": "Encrypted file successfully sent to the client"})
}
