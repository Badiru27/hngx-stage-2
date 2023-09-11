package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectToDB() *mongo.Client {

	clientOption := options.Client().ApplyURI(EnvMongoURI())

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, error := mongo.Connect(ctx, clientOption)

	if error != nil {
		log.Fatal(error)
	}

	error = client.Ping(ctx, nil)

	if error != nil {
		log.Fatal(error)
	}

	fmt.Println("Connected to MonogoDB successfully")

	return client

}

//Client instance
var DB *mongo.Client = ConnectToDB()

// function to get database collection
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("hngXStage2").Collection(collectionName)
	return collection
}
