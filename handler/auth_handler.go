package handler

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/TPM-Project-Larces/back-end.git/config"
	"github.com/TPM-Project-Larces/back-end.git/model"
	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

var secretKey = "XtgVuGrnPZksoM7FwBNSTVsKacjbp/Lx7zVUon9g/ls="

// @BasePath /

// @Summary Create user
// @Description Create a new user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body schemas.LoginRequest true "Request body"
// @Success 200 {object} schemas.AuthResponse
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /auth/login [post]
func Login(ctx *gin.Context) {

	request := schemas.LoginRequest{}
	ctx.BindJSON(&request)

	filter := bson.M{"email": request.Email, "password": request.Password}

	var loginUser model.User
	userCollection := config.GetMongoDB().Collection("user")
	err := userCollection.FindOne(context.Background(), filter).Decode(&loginUser)
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

func generateBearerToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	})

	// When the Docker container is created, set this secret key as an environment variable
	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "generate_token_error", err
	}
	tokenString = "Bearer " + tokenString
	return tokenString, nil
}

func findUserByUsername(username string) (string, error) {
	filter := bson.M{"username": username}

	var loginUser model.User
	userCollection := config.GetMongoDB().Collection("user")
	err := userCollection.FindOne(context.Background(), filter).Decode(&loginUser)
	if err != nil {
		return username, nil
	}

	return "", fmt.Errorf("user_not_found")
}

func MiddlewaveVerifyToken(ctx *gin.Context) (string, error) {
	tokenValue := ctx.GetHeader("Authorization")

	username, err := verifyToken(tokenValue)

	if err != nil {
		return "", err
	}

	username, err = findUserByUsername(username)

	if err != nil {
		return "", err
	}
	
	fmt.Printf("deu bom")
	return username, nil
}

func verifyToken(tokenValue string) (string, error) {
	token, err := jwt.Parse(removeBearerPrefix(tokenValue), func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); ok {
			return []byte(secretKey), nil
		}

		err := fmt.Errorf("invalid_token")
		return nil, err
	})

	if err != nil {
		err := fmt.Errorf("invalid_token")
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		err := fmt.Errorf("invalid_token")
		return "", err
	}
}

func removeBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		token = strings.TrimPrefix(token, "Bearer ")
	}

	return token
}
