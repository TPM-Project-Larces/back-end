package handler

import (
	"github.com/TPM-Project-Larces/back-end.git/schemas"
)

type CreateUserRequest struct {
	CPF         string          `bson:"cpf"`
	Name        string          `bson:"name"`
	Username    string          `bson:"username"`
	DateOfBirth string          `bson:"date_of_birth"`
	Email       string          `bson:"email"`
	Password    string          `bson:"password"`
	Contact     schemas.Contact `bson:"contact"`
	Address     schemas.Address `bson:"address"`
}

type DeleteUserRequest struct {
	Username string `bson:"username"`
}

type UpdateUserRequest struct {
	CPF         string          `bson:"cpf"`
	Name        string          `bson:"name"`
	Username    string          `bson:"username"`
	DateOfBirth string          `bson:"date_of_birth"`
	Email       string          `bson:"email"`
	Password    string          `bson:"password"`
	Contact     schemas.Contact `bson:"contact"`
	Address     schemas.Address `bson:"address"`
}

type LoginRequest struct {
	Email       string          `bson:"email"`
	Password    string          `bson:"password"`
}
