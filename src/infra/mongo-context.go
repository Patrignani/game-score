package infra

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoContext interface {
	GetCollection(database, collection string) *mongo.Collection
}

type MongoContextImp struct {
	connection string
}

func NewMongoContext(connectionString string) MongoContext {
	return &MongoContextImp{connectionString}
}

var optionsCredential = options.Credential{
	Username: "root",
	Password: "rootpassword",
}

func (m MongoContextImp) GetCollection(database, collection string) *mongo.Collection {
	client := m.getConnection()
	return client.Database(database).Collection(collection)
}

func (m MongoContextImp) getOptions() *options.ClientOptions {
	return options.Client().ApplyURI(m.connection).SetAuth(optionsCredential)
}

func (m MongoContextImp) getConnection() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), m.getOptions())

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	return client
}
