package handler

import "github.com/TPM-Project-Larces/back-end.git/schemas"

type LoginResponse struct {
	Message string        `bson:"message"`
	Data    schemas.Token `bson:"data"`
}

type CreateUserResponse struct {
	Message string               `bson:"message"`
	Data    schemas.UserResponse `bson:"data"`
}

type DeleteUserResponse struct {
	Message string               `bson:"message"`
	Data    schemas.UserResponse `bson:"data"`
}
type ShowUserResponse struct {
	Message string               `bson:"message"`
	Data    schemas.UserResponse `bson:"data"`
}
type ListUsersResponse struct {
	Message string                 `bson:"message"`
	Data    []schemas.UserResponse `bson:"data"`
}
type UpdateUserResponse struct {
	Message string               `bson:"message"`
	Data    schemas.UserResponse `bson:"data"`
}
