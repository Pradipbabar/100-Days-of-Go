package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Person represents a MongoDB document
type Person struct {
	Name    string `bson:"name"`
	Age     int    `bson:"age"`
	Address string `bson:"address"`
}

func main() {
	// Set up MongoDB connection
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	// Accessing the "people" collection in the "test" database
	collection := client.Database("test").Collection("people")

	// Insert a document
	person := Person{Name: "John Doe", Age: 30, Address: "123 Main St"}
	insertResult, err := collection.InsertOne(ctx, person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Inserted document with ID: %v\n", insertResult.InsertedID)

	// Find a document
	var result Person
	err = collection.FindOne(ctx, bson.M{"name": "John Doe"}).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found document: %+v\n", result)

	// Update a document
	updateResult, err := collection.UpdateOne(
		ctx,
		bson.M{"name": "John Doe"},
		bson.D{
			{"$set", bson.D{{"age", 31}}},
		},
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v document and modified %v document(s)\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Aggregation example: Find average age
	pipeline := bson.A{
		bson.D{{"$group", bson.D{{"_id", nil}, {"averageAge", bson.D{{"$avg", "$age"}}}}}},
	}
	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	var resultAggregate bson.M
	if cursor.Next(ctx) {
		err := cursor.Decode(&resultAggregate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Average Age: %v\n", resultAggregate["averageAge"])
	}
}
