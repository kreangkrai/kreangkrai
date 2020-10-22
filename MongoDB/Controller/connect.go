package Controller

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	collection *mongo.Collection
	client     *mongo.Client
	err        error
	ctx        context.Context
	cancel     context.CancelFunc
)

func Connect() {
	// Replace the uri string with your MongoDB deployment's connection string.
	uri := "mongodb+srv://Meeci:Meeci50026@meego.biqun.mongodb.net/Mee?retryWrites=true&w=majority"
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	// defer func() {
	// 	if err = client.Disconnect(ctx); err != nil {
	// 		panic(err)
	// 	}
	// }()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}
	//fmt.Println("Successfully connected and pinged.")

	collection = client.Database("Mee").Collection("mee_farm")

}

func Disconnect() {
	client.Disconnect(ctx)
}
