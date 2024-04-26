package controller

import (
	"context"
	"fmt"
	"log"

	model "github.com/yatendrayadav2832/mongoapi/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://yatendrayadav2832:yatendra@cluster0.ldyncxi.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
const dbname = "netflix"
const collectionname = "watchlist"

var collection *mongo.Collection

func init() {
	// Client options
	clientOption := options.Client().ApplyURI(connectionString)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")
	collection = client.Database(dbname).Collection(collectionname)

	// collection instance
	fmt.Println("Collection instance is ready")

}

// insert 1 records
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted 1 Movie in db with the id:", inserted.InsertedID)
}

// Update 1 record
func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified Count : ", result.ModifiedCount)
}

// delete 1 Record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie Got delete with the deleteCount", deleteCount)

}

func deleteAllMovie() int64 {
	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Number of Movies Deletes : ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}

// Read Movies data
func getAllMovies() []primitive.M {
	// Find all documents in the collection
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	// Iterate over the cursor and decode each document into a map
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		// Append the decoded document (movie) to the movies slice
		movies = append(movies, movie)

	}
	// Close the cursor
	defer cur.Close(context.Background())

	// Return the slice of movies
	return movies
}
