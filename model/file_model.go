package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type EncryptedFile struct {
	ID               primitive.ObjectID `bson:"_id,omitempty"`
	Username         string             `bson:"username"`
	Name             string             `bson:"name"`
	Data             []byte             `bson:"data"`
	LocallyEncrypted bool               `bson:"locally_encrypted"`
	Size             int                `bson:"size"`
	//AnonymizedFile AnonymizedFile     `bson:"anonymized_file"`
	CreatedAt time.Time `bson:"created_at"`
}

type StringData struct {
	Data string `json:"data"`
}

/*type AnonymizedFile struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	Data                    []byte             `bson:"data"`
	AnonymizationTechinique string             `bson:"anonymization_techinique"`
	CreatedAt               time.Time          `bson:"created_at"`
}*/
