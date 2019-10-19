package repository

import (
	"../entity"

	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

func InsertNewPerson(person entity.Person) interface{} {
	client := GetClient()
	collection := client.Database("civilact").Collection("person")
	insertResult, err := collection.InsertOne(context.TODO(), person)
	if err != nil {
		log.Fatalln("Error on inserting new Person", err)
	}
	return insertResult.InsertedID
}

func GetClient() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
