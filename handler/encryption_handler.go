package handler

import (
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

// @BasePath /
// @Summary Upload key
// @Description Uploads a public key
// @Tags Encryption
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "key_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /encryption/upload_key [post]
func UploadKey(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	err = ctx.SaveUploadedFile(file, "./key/"+file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	response(ctx, 200, "keys_generated", nil)
}

// @BasePath /
// @Summary Decrypt a file
// @Description Provide the filename to decrypt
// @Tags Encryption
// @Produce json
// @Param filename formData string true "Filename to decrypt"
// @Success 200 {string} string "file_decrypted"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /encryption/decrypt_file [post]
func DecryptFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	nameFile := ctx.PostForm("filename")

	uploadDir := "./encrypted_files"
	filePath := filepath.Join(uploadDir, nameFile)
	_, err := os.Stat(filePath)

	if nameFile == "" || os.IsNotExist(err) {
		response(ctx, 404, "file_not_found", nil)
		return
	} else if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	url := "http://localhost:3000/encryption/decrypt_file/"
	if err := sendFile(filePath, url); err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	response(ctx, 200, "file_decrypted", nil)
}
