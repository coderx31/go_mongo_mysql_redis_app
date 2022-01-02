package mongolib

import (
	"context"
	"fmt"
	"log"
	"mongo/configs"
	"mongo/models"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDataStore struct {
	Config  *configs.Mongo
	DB      *mongo.Database
	Session *mongo.Client
}

// creating new data store

func NewMongoDataStore(mongoConfigs *configs.Mongo) (*MongoDataStore, error) {
	mongoDataStore := new(MongoDataStore)
	mongoDataStore.Config = mongoConfigs
	database, client, err := connect(mongoConfigs)

	if database != nil && client != nil {
		mongoDataStore.DB = database
		mongoDataStore.Session = client
	}

	return mongoDataStore, err
}

func connect(mongoConfigs *configs.Mongo) (*mongo.Database, *mongo.Client, error) {
	var connectOnce sync.Once
	var database *mongo.Database
	var client *mongo.Client
	var err error

	connectOnce.Do(func() {
		database, client, err = connectToMongo(mongoConfigs)
	})

	return database, client, err

}

func connectToMongo(mongoConfigs *configs.Mongo) (*mongo.Database, *mongo.Client, error) {
	connStr := mongoConfigs.URI
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connStr))

	if err != nil {
		log.Fatal(err)
	}

	DB := client.Database(mongoConfigs.DB)

	return DB, client, nil

}

// get all people
func GetPeople(ctx context.Context, collection *mongo.Collection) (*[]models.Person, error) {
	cursor, err := collection.Find(ctx, bson.D{})

	if err != nil {
		fmt.Println(err.Error())
		defer cursor.Close(ctx)
		return nil, err
	}

	var people []models.Person
	if err = cursor.All(ctx, &people); err != nil {
		return nil, err
	}

	return &people, nil
}

// get specific person
func GetPerson(ctx context.Context, collection *mongo.Collection, firstname string) (*models.Person, error) {
	var person models.Person

	err := collection.FindOne(ctx, bson.D{{"firstname", firstname}}).Decode(&person)

	if err != nil {
		return nil, err
	}

	return &person, nil
}

// insert person
func CreatePerson(ctx context.Context, collection *mongo.Collection, person models.Person) (*models.Person, error) {
	result, err := collection.InsertOne(ctx, &person)

	if err != nil {
		return nil, err
	}
	fmt.Println(result)

	return &person, nil
}

// update person
func UpdatePerson(ctx context.Context, collection *mongo.Collection, firstname string, age int) (interface{}, error) {
	filter := bson.M{"firstname": firstname}
	update := bson.D{{"$set", bson.D{{"age", age}}}} // remember commas !!! not sem-colons
	result, err := collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return nil, err
	}
	//fmt.Println(result)

	return result, nil
}

func DeletePerson(ctx context.Context, collection *mongo.Collection, firstname string) (interface{}, error) {
	filter := bson.D{{"firstname", firstname}}
	result, err := collection.DeleteOne(ctx, filter)

	if err != nil {
		return nil, err
	}

	return result, err
}
