package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublicKey struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	KeyBytes  []byte             `bson:"key_bytes"`
	CreatedAt time.Time          `bson:"created_at"`
}
