package handler

import (
	"context"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /

// @Summary Create user
// @Description Create a new user
// @Tags Register
// @Accept json
// @Produce json
// @Param request body CreateUserRequest true "Request body"
// @Success 200 {object} CreateUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /create_user [post]
func CreateUser(ctx *gin.Context) {

	request := CreateUserRequest{}

	ctx.BindJSON(&request)

	// Acesse a coleção "users"
	collection := db.Collection("users")

	// Cria uma instância da struct User com os dados a serem inseridos
	user := schemas.User{
		CPF:         request.CPF,
		Name:        request.Name,
		DateOfBirth: request.DateOfBirth,
		Email:       request.Email,
		Password:    request.Password,
		Contact:     request.Contact,
		Address:     request.Address,
	}

	// Insira a struct na coleção
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		response(ctx, 400, "bad_request", err)
	}

	response(ctx, 200, "user_created", err)
}
