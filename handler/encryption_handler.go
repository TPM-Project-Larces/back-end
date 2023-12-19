package handler

import (
	"context"
	"github.com/TPM-Project-Larces/back-end.git/config"
	"github.com/TPM-Project-Larces/back-end.git/model"
	"go.mongodb.org/mongo-driver/bson"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

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
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", nil)
		return
	}

	ctx.Request.ParseMultipartForm(10 << 20)

	file, err := ctx.FormFile("arquivo")
	if err != nil {
		response(ctx, 400, "bad_request", nil)
		return
	}

	tempDir := "./key"
	err = os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	tempFile, err := os.Create(tempDir + "/" + file.Filename)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}
	defer tempFile.Close()

	err = ctx.SaveUploadedFile(file, tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	publicKeyPEM, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	keyCollection := config.GetMongoDB().Collection("key")

	// Checks if the user has a key in database
	existingKey := model.PublicKey{}
	err = keyCollection.FindOne(context.Background(), bson.M{"username": username}).Decode(&existingKey)
	if err == nil {
		// If the user already has a key, delete the previous key
		_, err := keyCollection.DeleteOne(context.Background(), bson.M{"username": username})
		if err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
	}

	key := model.PublicKey{
		Username:  username,
		KeyBytes:  publicKeyPEM,
		CreatedAt: time.Now(),
	}

	_, err = keyCollection.InsertOne(context.Background(), key)
	if err != nil {
		response(ctx, 500, "key_not_created", err)
		return
	}

	err = os.Remove(tempFile.Name())
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	ctx.JSON(200, gin.H{"message": "key_uploaded", "username": username})
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
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil || username == "" {
		response(ctx, 403, "invalid_token", nil)
		return
	}

	ctx.Request.ParseMultipartForm(10 << 20)

	nameFile := ctx.PostForm("filename")

	uploadDir := "./encrypted_files"
	filePath := filepath.Join(uploadDir, nameFile)
	_, err = os.Stat(filePath)

	if nameFile == "" || os.IsNotExist(err) {
		response(ctx, 404, "file_not_found", nil)
		return
	} else if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	collection := config.GetMongoDB().Collection("files")

	cursor, err := collection.Find(ctx, bson.M{"username": username, "name": nameFile})
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer cursor.Close(ctx)

	if cursor.Next(ctx) {
		var encryptedFile model.EncryptedFile
		if err := cursor.Decode(&encryptedFile); err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}

		tempFile, err := ioutil.TempFile("", "encrypted_file_*.bin")
		if err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
		defer tempFile.Close()

		if _, err := tempFile.Write(encryptedFile.Data); err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}

		if UploadAttestation(ctx) != nil {
			response(ctx, 500, "attestation_failed", nil)
			return
		}

		url := "http://localhost:3000/encryption/decrypt_file/"
		if err := sendFile(filePath, url); err != nil {
			response(ctx, 500, "internal_server_error", nil)
			return
		}

		if err := os.Remove(tempFile.Name()); err != nil {
			response(ctx, 500, "internal_server_error", nil)
		}

		response(ctx, 200, "file_decrypted", nil)

	} else if err := cursor.Err(); err != nil {
		response(ctx, 500, "internal_server_error", nil)
	}
}
