package main

import (
	"context"
	"log"
	"math/rand"
	"mongodemo/airbnb"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func waitForSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}

func main() {
	ctx := context.Background()

	client, err := airbnb.ConnectToMongo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	airbnb.QueryAirBnB(ctx, client)
	go airbnb.WatchForInserts(ctx, client)
	go airbnb.CreateListingsAndReviewsForever(ctx, client)

	waitForSignal()
}
