package handler

import (
	"context"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Delete file
// @Description deletes a file
// @Tags Files
// @Produce json
// @Param request body DeleteFileRequest true "Request body"
// @Success 200 {object} DeleteFileResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /files [delete]
func DeleteFile(ctx *gin.Context) {
	request := DeleteFileRequest{}
	ctx.BindJSON(&request)

	collection := db.Collection("files")

	file := DeleteFileRequest{

		Filename: request.Filename,
	}

	filter := bson.M{"name": file.Filename}

	// Busque o usuário antes de excluí-lo
	var deletedFile schemas.EncryptedFile
	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&deletedFile)

	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	// Se o usuário foi encontrado e excluído com sucesso, retorne os detalhes do usuário excluído
	ctx.JSON(200, gin.H{"message": "file_deleted", "deletedFile": deletedFile})
}
