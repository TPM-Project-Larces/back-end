package handler

import (
	"context"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Get all encrypted files
// @Description Get a list of all encrypted files
// @Tags Files
// @Accept json
// @Produce json
// @Success 200 {object} ListFilesResponse
// @Failure 500 {string} string "internal_server_error"
// @Router /files [get]
func GetFiles(ctx *gin.Context) {
	// Acesse a coleção "file"
	collection := db.Collection("files")

	// Execute uma consulta para buscar todos os documentos na coleção
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer cursor.Close(ctx)

	// Crie uma slice para armazenar os resultados
	var files []schemas.EncryptedFile

	// Itere sobre os resultados e decodifique cada documento em uma estrutura de arquivo
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
