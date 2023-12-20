package handler

import (
	"context"

	"github.com/TPM-Project-Larces/back-end.git/config"
	"github.com/TPM-Project-Larces/back-end.git/model"
	"github.com/TPM-Project-Larces/back-end.git/schemas"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

// @BasePath /
// @Summary Get all users
// @Description Get a list of all users
// @Tags User
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer JWT token" @in header
// @Success 200 {object} schemas.ListUsersResponse
// @Failure 500 {string} string "internal_server_error"
// @Router /user [get]
func GetUsers(ctx *gin.Context) {
	_, err := MiddlewaveVerifyToken(ctx)
	if err != nil {
		response(ctx, 403, "invalid_token", err)
		return
	}

	userCollection := config.GetMongoDB().Collection("user")

	search, err := userCollection.Find(ctx, bson.M{})
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}
	defer search.Close(ctx)

	var users []model.User

	for search.Next(context.Background()) {
		var user model.User

		if err := search.Decode(&user); err != nil {
			response(ctx, 500, "internal_server_error", err)
			return
		}

		users = append(users, user)
	}

	ctx.JSON(200, gin.H{"users": users})
}

// @BasePath /
// @Summary Find user by username
// @Description Provide the user data
// @Tags User
// @Produce json
// @Param username query string true "User`s username to find"
// @Success 200 {object} schemas.ShowUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /user/username [get]
func GetUserByUsername(ctx *gin.Context) {
	_, err := MiddlewaveVerifyToken(ctx)
	if err != nil {
		response(ctx, 403, "invalid_token", err)
		return
	}

	username := ctx.Query("username")

	userCollection := config.GetMongoDB().Collection("user")

	filter := bson.M{"username": username}

	var result model.User
	err = userCollection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		response(ctx, 400, "bad_request", err)
		return
	}

	ctx.JSON(200, gin.H{"user": result})

}

// @BasePath /
// @Summary Create user
// @Description Create a new user
// @Tags User
// @Accept json
// @Produce json
// @Param request body schemas.CreateUserRequest true "Request body"
// @Success 200 {object} schemas.CreateUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 500 {string} string "internal_server_error"
// @Router /user [post]
func CreateUser(ctx *gin.Context) {

	request := schemas.CreateUserRequest{}
	ctx.BindJSON(&request)

	userCollection := config.GetMongoDB().Collection("user")

	user := model.User{
		CPF:         request.CPF,
		Name:        request.Name,
		Username:    request.Username,
		DateOfBirth: request.DateOfBirth,
		Email:       request.Email,
		Password:    request.Password,
		Contact:     request.Contact,
		Address:     request.Address,
	}

	err := userCollection.FindOne(context.Background(),
		bson.M{"username": user.Username}).Decode(&user)
	if err == nil {
		response(ctx, 500, "username_already_exists", err)
		return
	}

	err = userCollection.FindOne(context.Background(),
		bson.M{"email": user.Email}).Decode(&user)
	if err == nil {
		response(ctx, 500, "email_already_exists", err)
		return
	}

	_, err = userCollection.InsertOne(context.Background(), user)
	if err != nil {
		response(ctx, 500, "user_not_created", err)
	}

	response(ctx, 200, "user_created", err)
}

// @BasePath /
// @Summary Update user
// @Description Updates a user
// @Tags User
// @Produce json
// @Param username query string true "User's username"
// @Param user body schemas.UpdateUserRequest true "User data to Update"
// @Success 200 {object} schemas.UpdateUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "user_not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /user [put]
func UpdateUser(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil {
		response(ctx, 403, "invalid_token", err)
		return
	}

	request := schemas.UpdateUserRequest{}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userCollection := config.GetMongoDB().Collection("user")

	filter := bson.M{"username": username}

	update := bson.M{
		"$set": request,
	}

	result, err := userCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if result.ModifiedCount == 0 {
		ctx.JSON(404, gin.H{"message": "user_not_found"})
		return
	}

	ctx.JSON(200, gin.H{"message": "user_updated"})
}

// @BasePath /
// @Summary Delete user
// @Description Deletes a user
// @Tags User
// @Produce json
// @Param request body schemas.DeleteUserRequest true "Request body"
// @Success 200 {object} schemas.DeleteUserResponse
// @Failure 400 {string} string "bad_request"
// @Failure 404 {string} string "not_found"
// @Failure 500 {string} string "internal_server_error"
// @Router /user [delete]
func DeleteUser(ctx *gin.Context) {
	username, err := MiddlewaveVerifyToken(ctx)
	if err != nil {
		response(ctx, 403, "invalid_token", err)
		return
	}

	request := schemas.DeleteUserRequest{}
	ctx.BindJSON(&request)

	collection := config.GetMongoDB().Collection("user")

	var deletedUser model.User
	filter := bson.M{"username": username}

	err = collection.FindOneAndDelete(context.Background(), filter).Decode(&deletedUser)
	if err != nil {
		response(ctx, 500, "internal_server_error", err)
		return
	}

	ctx.JSON(200, gin.H{"message": "user_deleted", "deletedUser": deletedUser})
}
