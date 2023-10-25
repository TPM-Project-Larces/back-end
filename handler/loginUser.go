package handler

import (
	"context"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func generateBearerToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
	})

	// When the Docker container is created, set this secret key as an environment variable
	tokenString, err := token.SignedString([]byte("XtgVuGrnPZksoM7FwBNSTVsKacjbp/Lx7zVUon9g/ls="))
	if err != nil {
		return "generate_token_error", err
	}

	return tokenString, nil
}

// @BasePath /

// @Summary Create user
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "Request body"
// @Success 200 {object} LoginResponse
// @Failure 404 {string} string "user_not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /login [post]
func Login(ctx *gin.Context) {

	request := LoginRequest{}
	ctx.BindJSON(&request)

	filter := bson.M{"email": request.Email, "password": request.Password}

	var loginUser schemas.User
	collection := db.Collection("users")
	err := collection.FindOne(context.Background(), filter).Decode(&loginUser)
	if err != nil {
		response(ctx, 404, "user_not_found", nil)
		return
	}

	token, err := generateBearerToken(loginUser.Username)
	if err != nil {
		response(ctx, 500, "internal_server_error", nil)
		return
	}

	ctx.JSON(200, gin.H{"message": "user_logged", "token": token})
}
