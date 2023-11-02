package handler

import (
	"context"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Get encrypted files by username
// @Description Get a list of encrypted files by username
// @Tags Files
// @Accept json
// @Produce json
// @Param username query string true "Username"
// @Success 200 {object} ListFilesResponse
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /files/by_username [get]
func GetFilesByUsername(ctx *gin.Context) {

	username := ctx.Query("username")
	if username == "" {
		response(ctx, 400, "Username parameter is required", nil)
		return
	}

	collection := db.Collection("files")

	cursor, err := collection.Find(ctx, bson.M{"username": username})
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer cursor.Close(ctx)

	var files []schemas.EncryptedFile

	for cursor.Next(context.Background()) {
		var file schemas.EncryptedFile
		if err := cursor.Decode(&file); err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
		files = append(files, file)
	}

	ctx.JSON(200, gin.H{"files": files})
}
