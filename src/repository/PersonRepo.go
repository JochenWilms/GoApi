package repository

import (
	"../entity"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

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
	return getMultiplePersonByQuery(bson.M{})
}

func GetPersonByName(firstName string) entity.Person {
	return getPersonByQuery(bson.M{"firstname": firstName})
}

func FindPersonByMail(email string) entity.Person {
	return getPersonByQuery(bson.M{"email": email})
}

func UpdatePassword(id string, password string) bool {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		fmt.Println(err)
	}
	result := updatePersonByQuery(bson.M{"_id": objectId}, bson.M{
		"$set": bson.M{
			"password": password,
		},
	})
	fmt.Println(result)
	if result.ModifiedCount == 1 {
		return true
	}
	return false
}

func updatePersonByQuery(filter bson.M, update bson.M) *mongo.UpdateResult {
	client := getClient()
	collection := client.Database(database).Collection(collection)

	result, err := collection.UpdateOne(
		context.TODO(),
		filter,
		update,
	)
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func getMultiplePersonByQuery(query bson.M) []entity.Person {
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

func getPersonByQuery(query bson.M) entity.Person {
	client := getClient()
	collection := client.Database(database).Collection(collection)
	//Set the limit of the number of record to find
	//Define an array in which you can store the decoded documents
	var result entity.Person

	//Passing the bson.D{{}} as the filter matches  documents in the collection
	err := collection.FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	return result
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
