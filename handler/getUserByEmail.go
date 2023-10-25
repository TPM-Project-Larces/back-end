package handler

import (
	"context"
	"fmt"

	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Find user by username
// @Description Provide the user data
// @Tags Users
// @Produce json
// @Param username query string true "User`s username to find"
// @Success 200 {object} ShowUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /user/by_username [get]
func GetUserByUsername(ctx *gin.Context) {

	username := ctx.Query("username")
	fmt.Println("username: " + username)
	collection := db.Collection("users")

	filter := bson.M{"username": username}

	var result schemas.User

	err := collection.FindOne(context.Background(), filter).Decode(&result)

	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	ctx.JSON(200, gin.H{"user": result})

}
