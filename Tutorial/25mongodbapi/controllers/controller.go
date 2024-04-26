package controller

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://yatendrayadav2832:yatendra@cluster0.ldyncxi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbname = "netflix"
const collectionname = "watchlist"

var collection *mongo.Collection

func init()  {
	// Client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client,err := mongo.Connect(context.TODO(),clientOption)
	if err!=nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")
	collection = client.Database(dbname).Collection(collectionname)

	// collection instance
	fmt.Println("Collection instance is ready")

}