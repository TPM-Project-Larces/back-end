package handler

import (
	"context"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Delete user
// @Description deletes a user
// @Tags Users
// @Produce json
// @Param request body DeleteUserRequest true "Request body"
// @Success 200 {object} DeleteUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /user [delete]
func DeleteUser(ctx *gin.Context) {
	request := DeleteUserRequest{}
	ctx.BindJSON(&request)

	collection := db.Collection("users")

	username := DeleteUserRequest{

		Username: request.Username,
	}

	// Crie um filtro para encontrar o usuário pelo email
	filter := bson.M{"username": username.Username}

	// Busque o usuário antes de excluí-lo
	var deletedUser schemas.User
	err := collection.FindOneAndDelete(context.Background(), filter).Decode(&deletedUser)

	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	// Se o usuário foi encontrado e excluído com sucesso, retorne os detalhes do usuário excluído
	ctx.JSON(200, gin.H{"message": "user_deleted", "deletedUser": deletedUser})
}
