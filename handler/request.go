package handler

import "github.com/TPM-Project-Larces/back-end.git/schemas"

type CreateUserRequest struct {
	CPF         string          `bson:"cpf"`
	Name        string          `bson:"name"`
	DateOfBirth string          `bson:"date_of_birth"`
	Email       string          `bson:"email"`
	Password    string          `bson:"password"`
	Contact     schemas.Contact `bson:"contact"`
	Address     schemas.Address `bson:"address"`
}
