package handler

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
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
// @Summary Upload file
// @Description Upload a file to encrypt
// @Tags Encryption
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "file_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /encryption/upload_file [post]
func UploadFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	// Open public key file
	filePath := "./key/public_key.pem"
	filePublicKey, err := os.Open(filePath)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer filePublicKey.Close()

	// Reads public key file
	publicKeyData, err := ioutil.ReadAll(filePublicKey)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	blockPublicKey, _ := pem.Decode(publicKeyData)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	publicKeyData = blockPublicKey.Bytes

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyData)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	publicKeyRsa := publicKey.(*rsa.PublicKey)

	// Open the file directly without saving it to disk
	uploadedFile, err := file.Open()
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer uploadedFile.Close()

	data, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	// Split data into smaller blocks (maximum block size for RSA encryption)
	maxBlockSize := 245
	var encryptedBlocks []byte
	for len(data) > 0 {
		blockSize := len(data)
		if blockSize > maxBlockSize {
			blockSize = maxBlockSize
		}

		// Encrypt the block and add to the list of encrypted blocks
		encryptedBlock, err := rsa.EncryptPKCS1v15(rand.Reader, publicKeyRsa, data[:blockSize])
		if err != nil {
			response(ctx, 500, "internal_server_error", nil)
			return
		}
		encryptedBlocks = append(encryptedBlocks, encryptedBlock...)
		data = data[blockSize:]
	}

	tempDir := "./encrypted_files"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	tempfile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer tempfile.Close()

	err = ioutil.WriteFile(tempfile.Name(), encryptedBlocks, 0644)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	response(ctx, 200, "file_uploaded", nil)
}

// @BasePath /
// @Summary Save file
// @Description Save a file to encrypt
// @Tags Encryption
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "file_saved"
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /encryption/saved_file [post]
func SavedFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	// Open the file directly without saving it to disk
	uploadedFile, err := file.Open()
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer uploadedFile.Close()

	data, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	// Split data into smaller blocks
	maxBlockSize := 245
	var encryptedBlocks []byte
	for len(data) > 0 {
		blockSize := len(data)
		if blockSize > maxBlockSize {
			blockSize = maxBlockSize
		}

		// Write file in bytes blocks
		encryptedBlock := data[:blockSize]
		encryptedBlocks = append(encryptedBlocks, encryptedBlock...)
		data = data[blockSize:]
	}

	tempDir := "./decrypted_files"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	tempfile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer tempfile.Close()

	err = ioutil.WriteFile(tempfile.Name(), encryptedBlocks, 0644)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	response(ctx, 200, "file_saved", err)
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
