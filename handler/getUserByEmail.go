package handler

import (
	"context"
	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @Summary Find user database
// @Description Provide the user data
// @Tags Users
// @Produce json
// @Param email formData string true "User email to find"
// @Success 200 {object} ShowUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /get_user_by_email [post]
func GetUserByEmail(ctx *gin.Context) {
	emailUser := ctx.PostForm("emailUser")

	// Acesse a coleção "users"
	collection := db.Collection("users")

	filter := bson.D{{"email", emailUser}}
	result := schemas.User{}
	err := collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		response(ctx, 400, "bad_request", err)
	}

	ctx.JSON(200, gin.H{"user": result})
}
