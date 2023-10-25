package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Defining the Address model
type Address struct {
	Street  string
	City    string
	State   string
	ZIPCode string
}

// Defining the Contact model
type Contact struct {
	Celphone string
	Phone    string
}

// Defining the User model
type User struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CPF         string             `bson:"cpf"`
	Name        string             `bson:"name"`
	Username    string             `bson:"username"`
	DateOfBirth string             `bson:"date_of_birth"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Contact     Contact            `bson:"contact"`
	Address     Address            `bson:"address"`
}

type UserResponse struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt   time.Time          `bson:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at"`
	DeletedAt   *time.Time         `bson:"deleted_at,omitempty"`
	CPF         string             `bson:"cpf"`
	Name        string             `bson:"name"`
	Username    string             `bson:"username"`
	DateOfBirth string             `bson:"date_of_birth"`
	Email       string             `bson:"email"`
	Password    string             `bson:"password"`
	Contact     Contact            `bson:"contact"`
	Address     Address            `bson:"address"`
}

type Token struct {
	token string `bson:"token"`
}
