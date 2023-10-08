package handler

import (
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
// @Router /saved_file [post]
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

	tempDir := "./decrypted_files"
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

	response(ctx, 200, "file_decrypted", err)
}
