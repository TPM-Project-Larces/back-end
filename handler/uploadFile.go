package handler

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"io/ioutil"
	"os"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @BasePath /

// @Summary upload file

// @Description upload a file to encrypt
// @Tags Server operations
// @Accept multipart/form-data
// @Produce json
// @Param arquivo formData file true "File"
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
		return
	}

	// Verificar se um arquivo com o mesmo nome já existe no banco de dados
	name := file.Filename
	collection := db.Collection("files")
	existingFile := &schemas.EncryptedFile{}
	err = collection.FindOne(context.Background(), bson.M{"name": name}).Decode(existingFile)

	if err == nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	// Abrir o arquivo de chave pública
	filePath := "./key/public_key.pem"
	filePublicKey, err := os.Open(filePath)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer filePublicKey.Close()

	// Ler o arquivo de chave pública
	publicKeyData, err := ioutil.ReadAll(filePublicKey)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	blockPublicKey, _ := pem.Decode(publicKeyData)
	if blockPublicKey == nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	publicKeyData = blockPublicKey.Bytes

	publicKey, err := x509.ParsePKIXPublicKey(publicKeyData)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	publicKeyRsa, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	// Abrir o arquivo diretamente sem salvá-lo no disco
	uploadedFile, err := file.Open()
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer uploadedFile.Close()

	data, err := ioutil.ReadAll(uploadedFile)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
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
			return
		}
		encryptedBlocks = append(encryptedBlocks, encryptedBlock...)
		data = data[blockSize:]
	}

	tempDir := "./encrypted_files"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	tempfile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer tempfile.Close()

	err = ioutil.WriteFile(tempfile.Name(), encryptedBlocks, 0644)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	// PostEncryptedFile
	{
		name := file.Filename
		data := encryptedBlocks
		collection := db.Collection("files")
		file := schemas.EncryptedFile{Username: "username", Name: name, Data: data, LocallyEncrypted: false}
		_, err := collection.InsertOne(context.Background(), file)
		if err != nil {
			response(ctx, 400, "bad_request", err)
			return
		}
		response(ctx, 200, "encrypted_file_created", nil)
	}

	response(ctx, 200, "file_uploaded", nil)
}
