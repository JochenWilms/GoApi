package repository

import (
	"../entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"

	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

const (
	host       = "localhost:27017"
	database   = "GoApi"
	username   = "username"
	password   = "password"
	collection = "person"
)

func InsertNewPerson(person entity.Person) interface{} {
	client := getClient()
	collection := client.Database(database).Collection(collection)
	insertResult, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatalln("Error on inserting new Person", err)
	}
	return insertResult.InsertedID
}

func GetAllPersons() []entity.Person {
	return getPersonByQuery(bson.M{})
}

func GetPersonByName(firstName string) []entity.Person {
	return getPersonByQuery(bson.M{"firstname": firstName})
}

func getPersonByQuery(query bson.M) []entity.Person {
	client := getClient()
	collection := client.Database(database).Collection(collection)
	findOptions := options.Find()
	//Set the limit of the number of record to find
	//Define an array in which you can store the decoded documents
	var results []entity.Person

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	cur, err := collection.Find(context.TODO(), query, findOptions)
	if err != nil {
		log.Fatal(err)
	}
	//Finding multiple documents returns a cursor
	//Iterate through the cursor allows us to decode documents one at a time

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		var elem entity.Person
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	//Close the cursor once finished
	cur.Close(context.TODO())
	return results
}

func getClient() *mongo.Client {
	uri := fmt.Sprintf(`mongodb://%s:%s@%s/%s?authSource=admin`,
		username,
		password,
		host,
		database,
	)
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	return client
}
