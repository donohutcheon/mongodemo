package airbnb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func CreateListingsAndReviewsForever(ctx context.Context, client *mongo.Client) {
	for {
		insertListing(ctx, client)
		time.Sleep(1 * time.Second)
	}
}

// function to insert a random listing into the listingsAndReviews collection
func insertListing(ctx context.Context, client *mongo.Client) {
	collection := client.Database("sample_airbnb").Collection("listingsAndReviews")
	listing := randomListing()
	_, err := collection.InsertOne(ctx, listing)
	if err != nil {
		log.Fatal(err)
	}
}

func randomListing() *ListingAndReviews {
	return &ListingAndReviews{
		Name:               randomName(nameElements),
		MinimumNights:      strconv.Itoa(rand.Intn(4)),
		MaximumNights:      strconv.Itoa(rand.Intn(4) + 4),
		CancellationPolicy: "flexible",
		Accommodates:       rand.Intn(8),
		Bedrooms:           rand.Intn(8),
		Beds:               40,
		Bathrooms:          primitive.NewDecimal128(0, rand.Uint64()),
		Price:              primitive.NewDecimal128(0, rand.Uint64()),
	}
}
