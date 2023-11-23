package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PublicKeyResponse struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Data      []byte             `bson:"data"`
	CreatedAt time.Time          `bson:"created_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}
