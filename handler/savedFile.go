package handler

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary Upload encrypted file
// @Description upload a encrypted file
// @Tags Encryption
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "file_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /encryption/upload_encrypted_file [post]
func SavedFile(ctx *gin.Context) {
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", err)
	}

	//Abra o arquivo diretamente sem salvá-lo no disco
	uploadedFile, err := file.Open()
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}
	defer uploadedFile.Close()

	data, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
	}

	// Dividir os dados em blocos menores
	maxBlockSize := 245
	var encryptedBlocks []byte
	for len(data) > 0 {
		blockSize := len(data)
		if blockSize > maxBlockSize {
			blockSize = maxBlockSize
		}

		//Escreve o arquivo em blocos de bytes
		encryptedBlock := data[:blockSize]

		encryptedBlocks = append(encryptedBlocks, encryptedBlock...)
		data = data[blockSize:]
	}

	tempDir := "./locally_encrypted_files"
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
		response(ctx, 500, "internal_server_error", err)
	}

	// PostEncryptedFile
	{
		name := file.Filename
		data := encryptedBlocks
		collection := db.Collection("files")
		file := schemas.EncryptedFile{Username: "username", Name: name, Data: data, LocallyEncrypted: true}
		_, err := collection.InsertOne(context.Background(), file)
		if err != nil {
			response(ctx, 400, "bad_request", err)
			return
		}
		response(ctx, 200, "encrypted_file_created", nil)
	}

	response(ctx, 200, "file_decrypted", err)
}
