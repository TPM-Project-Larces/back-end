package handler

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// @Summary Decrypt a file
// @Description Provide the filename to decrypt
// @Tags Server operations
// @Produce json
// @Param filename formData string true "Filename to decrypt"
// @Success 200 {string} string "file_decrypted"
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
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
		response(ctx, 404, "file_not_found", err)
		return
	} else if err != nil {

		// Outro erro ocorreu, retorna um erro interno do servidor
		response(ctx, 500, "internal_server_error", err)
		return
	}

	url := "http://localhost:5000/decrypt_file/"

	sendFile(filePath, url)

	response(ctx, 200, "file_decrypted", err)

}
