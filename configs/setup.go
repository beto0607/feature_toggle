package configs

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func connectDB() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(EnvMongoURI()).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), 30*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	cancelFunc()
	return client
}

func GetDBClient() *mongo.Client {
	if databaseClient == nil {
		ConnectToDB()
	}
	return databaseClient
}

func ConnectToDB() {
	databaseClient = connectDB()
}

var databaseClient *mongo.Client

func GetCollection(collectionName string) *mongo.Collection {
	client := GetDBClient()
	collection := client.Database("toggler").Collection(collectionName)
	return collection
}
