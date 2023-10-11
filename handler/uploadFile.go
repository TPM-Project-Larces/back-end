package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary upload file

// @Description upload a file to encrypt
// @Tags Server operations
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "file_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /upload_file [post]
func UploadFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	// Open public key file
	filePath := "./key/public_key.pem"
	filePublicKey, err := os.Open(filePath)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}
	defer filePublicKey.Close()

	// Reads public key file
	publicKeyData, err := ioutil.ReadAll(filePublicKey)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	blockPublicKey, _ := pem.Decode(publicKeyData)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	publicKeyData = blockPublicKey.Bytes

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyData)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	publicKeyRsa := publicKey.(*rsa.PublicKey)

	// Abra o arquivo diretamente sem salvá-lo no disco
	uploadedFile, err := file.Open()
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}
	defer uploadedFile.Close()

	data, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	// Dividir os dados em blocos menores (tamanho máximo de bloco para criptografia RSA)
	maxBlockSize := 245
	var encryptedBlocks []byte
	for len(data) > 0 {
		blockSize := len(data)
		if blockSize > maxBlockSize {
			blockSize = maxBlockSize
		}

		// Criptografar o bloco e adicionar à lista de blocos criptografados
		encryptedBlock, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyRsa, data[:blockSize])
		if err != nil {
			response(ctx, 500, "internal_server_error", err)
		}
		encryptedBlocks = append(encryptedBlocks, encryptedBlock...)
		data = data[blockSize:]
	}

	tempDir := "./encrypted_files"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	tempfile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}
	defer tempfile.Close()

	err = ioutil.WriteFile(tempfile.Name(), encryptedBlocks, 0644)
	if err != nil {
		fmt.Println("Erro ao escrever o arquivo criptografado:", err)
		return
	}

	response(ctx, 200, "file_uploaded", err)

}
