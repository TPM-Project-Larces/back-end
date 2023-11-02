package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateMongoDBDatabase() (*mongo.Database, error) {
	// Define a configuração do cliente MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Conecte-se ao servidor MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {

		return nil, err
	}

	// Verifique a conexão com o MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	// Escolha ou crie um banco de dados
	database := client.Database("server_database")

	return database, nil
}
