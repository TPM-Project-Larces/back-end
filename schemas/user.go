package schemas

import "time"

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
	CPF         string
	Name        string
	DateOfBirth string
	Email       string
	Password    string
	Contact     Contact
	Address     Address
}

type UserResponse struct {
	ID          uint       `bson:"id"`
	CreatedAt   time.Time  `bson:"created_at"`
	UpdatedAt   time.Time  `bson:"updated_at"`
	DeletedAt   *time.Time `bson:"deleted_at,omitempty"`
	CPF         string     `bson:"cpf"`
	Name        string     `bson:"name"`
	DateOfBirth string     `bson:"date_of_birth"`
	Email       string     `bson:"email"`
	Password    string     `bson:"password"`
	Contact     Contact    `bson:"contact"`
	Address     Address    `bson:"address"`
}
