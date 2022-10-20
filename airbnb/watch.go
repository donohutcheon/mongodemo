package airbnb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

// WatchForInserts watches for inserts made to the listingsAndReviews collection.
func WatchForInserts(ctx context.Context, client *mongo.Client) {
	matchStage := bson.D{{"$match", bson.D{{"operationType", "insert"}}}}
	opts := options.ChangeStream().SetMaxAwaitTime(2 * time.Second)
	changeStream, err := client.Watch(
		ctx,
		mongo.Pipeline{matchStage},
		opts)
	if err != nil {
		log.Fatal(err)
	}

	defer changeStream.Close(ctx)
	for changeStream.Next(ctx) {
		fmt.Printf("Got a change: %v\n", changeStream.Current)
	}
}
