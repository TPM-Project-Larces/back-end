package schemas

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
	//AnonymizedFile AnonymizedFile     `bson:"anonymized_file"`
	CreatedAt time.Time `bson:"created_at"`
}

type EncryptedFileResponse struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Name     string             `bson:"name"`
	Data     []byte             `bson:"data"`
	//AnonymizedFile AnonymizedFile     `bson:"anonymized_file"`
	CreatedAt time.Time  `bson:"created_at"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty"`
}

/*type AnonymizedFile struct {
	ID                      primitive.ObjectID `bson:"_id,omitempty"`
	Data                    []byte             `bson:"data"`
	AnonymizationTechinique string             `bson:"anonymization_techinique"`
	CreatedAt               time.Time          `bson:"created_at"`
}*/
