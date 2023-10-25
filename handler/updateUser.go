package handler

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Update user
// @Description updates a user
// @Tags Users
// @Produce json
// @Param username query string true "User's username"
// @Param user body UpdateUserRequest true "User data to Update"
// @Success 200 {object} UpdateUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "user_not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	username := ctx.Query("username") // Obtenha o nome de usuário a ser atualizado a partir da consulta

	request := UpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	collection := db.Collection("users")

	// Crie um filtro com base no nome de usuário
	filter := bson.M{"username": username}

	// Crie uma operação de atualização com os dados fornecidos
	update := bson.M{
		"$set": request,
	}

	// Execute a operação de atualização
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		// Se nenhum documento for modificado, o usuário não foi encontrado
		ctx.JSON(404, gin.H{"message": "user_not_found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "User updated successfully"})
}
