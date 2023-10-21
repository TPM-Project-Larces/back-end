package handler

import (
	"context"
	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Get all users
// @Description Get a list of all users
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 {object} ListUsersResponse
// @Failure 500 {string} string "internal_server_error"
// @Router /get_users [get]
func GetUsers(ctx *gin.Context) {
	// Acesse a coleção "users"
	collection := db.Collection("users")

	// Execute uma consulta para buscar todos os documentos na coleção
	busca, err := collection.Find(ctx, bson.M{})
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer busca.Close(ctx)

	// Crie uma slice para armazenar os resultados
	var users []schemas.User

	// Itere sobre os resultados e decodifique cada documento em uma estrutura de usuário
	for busca.Next(context.Background()) {
		var user schemas.User
		if err := busca.Decode(&user); err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}
		users = append(users, user)
	}

	ctx.JSON(200, gin.H{"users": users})
}
