package airbnb

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func QueryAirBnB(ctx context.Context, client *mongo.Client) {
	collection := client.Database("sample_airbnb").Collection("listingsAndReviews")
	filter := bson.D{
		{
			"bedrooms", bson.D{
				{"$gte", 4},
			}},
		{
			"property_type", bson.D{
				{"$eq", "House"},
			}},
		{"$or", bson.A{
			bson.D{{"cancellation_policy", "flexible"}},
			bson.D{{"cancellation_policy", "moderate"}},
		}},
	}

	docCount, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(docCount)

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Fatal(err)
	}

	for cur.Next(ctx) {
		var listing ListingAndReviews
		err := cur.Decode(&listing)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("name : %s\n", listing.Name)
	}
}
