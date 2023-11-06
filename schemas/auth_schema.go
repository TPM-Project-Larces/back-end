package schemas

import (
	"github.com/TPM-Project-Larces/back-end.git/model"
)

type LoginResponse struct {
	Message string      `bson:"message"`
	Data    model.Token `bson:"data"`
}

type LoginRequest struct {
	Email    string `bson:"email"`
	Password string `bson:"password"`
}

type AuthResponse struct {
	Message string `bson:"message"`
}

type AuthRequest struct {
	Token model.Token `bson:"token"`
}
