package handler

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"github.com/TPM-Project-Larces/back-end.git/config"
	"github.com/TPM-Project-Larces/back-end.git/model"
	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// @BasePath /
// @Summary Search file
// @Description Search file
// @Tags Encryption
// @Accept json
// @Produce json
// @Param request body schemas.StringData true "Request body"
// @Success 200 {object} schemas.ShowFileResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /encryption/search_file [post]
func SearchFile(ctx *gin.Context) {
	request := schemas.StringData{}
	ctx.BindJSON(&request)

	name := schemas.StringData{
		Data: request.Data,
	}

	fmt.Println("aquiiii", name.Data)

	collection := config.GetMongoDB().Collection("files")

	filter := bson.M{"name": name.Data}

	var result model.EncryptedFile
	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			response(ctx, 404, "not_found", nil)
			return
		}
		// Lidar com outros erros possíveis
		response(ctx, 400, "bad_request", err)
		return
	}

	if len(result.Data) == 0 {
		response(ctx, 400, "bad_request", nil)
		return
	}

	// Verificar se o diretório "agent_files" existe e criar se não existir
	dir := "./agent_files"
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.Mkdir(dir, os.ModePerm)
		if err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
	}

	filePath := dir + "/" + name.Data
	err = ioutil.WriteFile(filePath, []byte(result.Data), 0644)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	fmt.Println("ate aq")
	url := "http://localhost:3000/encryption/save_file"
	if err := sendFile(filePath, url); err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	url2 := "http://localhost:3000/encryption/size_and_decrypt"
	if err := sendString(strconv.Itoa(result.Size), url2); err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	fmt.Println("ate aq22", result.Size)
	ctx.JSON(200, gin.H{"file_saved": filePath})
}

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

	destDirectory := "./key/"

	if err := os.MkdirAll(destDirectory, os.ModePerm); err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	destPath := filepath.Join(destDirectory, file.Filename)

	if err := ctx.SaveUploadedFile(file, destPath); err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	fileContent, err := ioutil.ReadFile(destPath)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	// Agora, fileContent contém o conteúdo do arquivo em bytes
	keyBytes := fileContent

	collection := config.GetMongoDB().Collection("key")
	key := model.PublicKey{Username: "username", KeyBytes: keyBytes}
	_, err = collection.InsertOne(context.Background(), key)
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	response(ctx, 200, "key_created", nil)
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

	fmt.Println("name: " + nameFile)
	collection := config.GetMongoDB().Collection("files")

	filter := bson.M{"name": nameFile}

	var result model.EncryptedFile

	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	fileContent := result.Data

	// Caminho do arquivo a ser descriptografado
	decryptedFilePath := "./" + result.Name

	err = ioutil.WriteFile(decryptedFilePath, fileContent, 0644)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	// Enviar o arquivo para outra API
	url := "http://localhost:3000/encryption/decrypt_file/"
	if err := sendFile(decryptedFilePath, url); err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	response(ctx, 200, "file_decrypted", nil)
}
