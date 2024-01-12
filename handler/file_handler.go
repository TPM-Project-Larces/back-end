package handler

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/TPM-Project-Larces/back-end.git/config"
	"github.com/TPM-Project-Larces/back-end.git/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @BasePath /
// @Summary Get all encrypted files
// @Description Get a list of all encrypted files
// @Tags File
// @Accept json
// @Produce json
// @Success 200 {object} schemas.ListFilesResponse
// @Failure 500 {string} string "internal_server_error"
// @Router /file [get]
func GetFiles(ctx *gin.Context) {
	_, err := MiddlewaveVerifyToken(ctx)
	if err != nil {
		response(ctx, 403, "invalid_token", err)
		return
	}

	if UploadAttestation(ctx) != nil {
		response(ctx, 500, "attestation_failed", nil)
		return
	}

	collection := config.GetMongoDB().Collection("files")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer cursor.Close(ctx)

	var files []model.EncryptedFile

	for cursor.Next(context.Background()) {
		var file model.EncryptedFile
		if err := cursor.Decode(&file); err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
		files = append(files, file)
	}

	ctx.JSON(200, gin.H{"files": files})
}

// @BasePath /
// @Summary Get encrypted files by username
// @Description Get a list of encrypted files by username
// @Tags File
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} schemas.ListFilesResponse
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /file/by_username [get]
func GetFilesByUsername(ctx *gin.Context) {
	_, err := MiddlewaveVerifyToken(ctx)
	if err != nil {
		response(ctx, 403, "invalid_token", err)
		return
	}

	username := ctx.Query("username")
	if username == "" {
		response(ctx, 400, "Username parameter is required", nil)
		return
	}

	collection := config.GetMongoDB().Collection("files")

	cursor, err := collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer cursor.Close(ctx)

	var files []model.EncryptedFile

	for cursor.Next(context.Background()) {
		var file model.EncryptedFile
		if err := cursor.Decode(&file); err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
		files = append(files, file)
	}

	ctx.JSON(200, gin.H{"files": files})
}

// @BasePath /
// @Summary Find file by name
// @Description Provide the file data
// @Tags File
// @Produce json
// @Param filename query string true "filename to find"
// @Success 200 {object} schemas.ShowFileResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /file/by_name [get]
func GetFileByName(ctx *gin.Context) {
	_, err := MiddlewaveVerifyToken(ctx)
	if err != nil {
		response(ctx, 403, "invalid_token", err)
		return
	}

	name := ctx.Query("filename")
	fmt.Println("name: " + name)
	collection := config.GetMongoDB().Collection("files")

	filter := bson.M{"name": name}

	var result model.EncryptedFile

	err = collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	ctx.JSON(200, gin.H{"file": result})

}

// @BasePath /
// @Summary Upload encrypted file
// @Description Upload a encrypted file
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "File"
// @Success 200 {string} string "file_saved"
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /file/upload_encrypted_file [post]
func SavedFile(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", err)
		return
	}
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	if UploadAttestation(ctx) != nil {
		response(ctx, 500, "attestation_failed", nil)
		return
	}

	// Check if a file with the same name already exists in the database
	collection := config.GetMongoDB().Collection("files")
	existingFile := &model.EncryptedFile{}
	err = collection.FindOne(context.Background(), bson.M{"name": file.Filename, "username": username}).Decode(existingFile)
	if err == nil {
		response(ctx, 400, "bad_request", err)
		return
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

	// Split data into smaller blocks
	maxBlockSize := 245
	var encryptedBlocks []byte
	for len(data) > 0 {
		blockSize := len(data)
		if blockSize > maxBlockSize {
			blockSize = maxBlockSize
		}

		//Writes the file in blocks of bytes
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

	{
		name := file.Filename
		data := encryptedBlocks
		file := model.EncryptedFile{Username: username, Name: name, Data: data, LocallyEncrypted: true}
		_, err := collection.InsertOne(context.Background(), file)
		if err != nil {
			response(ctx, 400, "bad_request", err)
			return
		}
		response(ctx, 200, "encrypted_file_created", nil)
	}

	response(ctx, 200, "file_saved", err)
}

// @BasePath /
// @Summary Upload file
// @Description Upload a file to encrypt
// @Tags File
// @Accept multipart/form-data
// @Produce json
// @Param arquivo formData file true "File"
// @Success 200 {string} string "file_uploaded"
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /file/upload_file [post]
func UploadFile(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", nil)
		return
	}
	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	if UploadAttestation(ctx) != nil {
		response(ctx, 500, "attestation_failed", nil)
		return
	}

	// Check if a file with the same name already exists in the database
	name := file.Filename
	collection := config.GetMongoDB().Collection("files")
	existingFile := &model.EncryptedFile{}
	err = collection.FindOne(context.Background(), bson.M{"name": name, "username": username}).Decode(existingFile)

	if err == nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	// Open public key
	key := &model.PublicKey{}
	keycollection := config.GetMongoDB().Collection("key")
	err = keycollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&key)
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	blockPublicKey, _ := pem.Decode(key.KeyBytes)
	if blockPublicKey == nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	publicKey, err := x509.ParsePKIXPublicKey(blockPublicKey.Bytes)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	publicKeyRsa, ok := publicKey.(*rsa.PublicKey)
	if !ok {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	// Open the file directly without saving it to disk
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

	// Split data into smaller blocks (maximum block size for RSA encryption)
	maxBlockSize := 245
	var encryptedBlocks []byte
	for len(data) > 0 {
		blockSize := len(data)
		if blockSize > maxBlockSize {
			blockSize = maxBlockSize
		}

		//Encrypt the block and add to the list of encrypted blocks
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

	{
		name := file.Filename
		data := encryptedBlocks
		collection := config.GetMongoDB().Collection("files")
		file := model.EncryptedFile{Username: username, Name: name, Data: data, LocallyEncrypted: false}
		_, err := collection.InsertOne(context.Background(), file)
		if err != nil {
			response(ctx, 400, "bad_request", err)
			return
		}
		response(ctx, 200, "encrypted_file_created", nil)
	}

	err = os.Remove(file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	response(ctx, 200, "file_uploaded", nil)
}

// @BasePath /
// @Summary Delete file
// @Description deletes a file
// @Tags File
// @Produce json
// @Param arquivo path string true "Nome do arquivo"
// @Success 200 {object} schemas.DeleteFileResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /file/{arquivo} [delete]
func DeleteFile(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", err)
		return
	}

	fileName := ctx.Param("arquivo")
	fmt.Println(fileName)

	if UploadAttestation(ctx) != nil {
		response(ctx, 500, "attestation_failed", nil)
		return
	}

	collection := config.GetMongoDB().Collection("files")

	filter := bson.M{"name": fileName, "username": username}

	// Search for the user before deleting them
	var deletedFile model.EncryptedFile
	err = collection.FindOneAndDelete(context.Background(), filter).Decode(&deletedFile)

	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	ctx.JSON(200, gin.H{"message": "file_deleted", "deletedFile": deletedFile})
}
