package handler

import (
	"context"
	"fmt"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Find file by name
// @Description Provide the file data
// @Tags Files
// @Produce json
// @Param filename query string true "filename to find"
// @Success 200 {object} ShowFileResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /files/by_name [get]
func GetFileByName(ctx *gin.Context) {

	name := ctx.Query("filename")
	fmt.Println("name: " + name)
	collection := db.Collection("files")

	filter := bson.M{"name": name}

	var result schemas.EncryptedFile

	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	ctx.JSON(200, gin.H{"file": result})

}
