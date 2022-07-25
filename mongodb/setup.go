package mongodb

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client   *mongo.Client
	Shippers *mongo.Collection
)

func Setup(uri string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().
		SetMaxPoolSize(200).
		SetMaxConnecting(20).
		ApplyURI(uri))
	Client = client
	Shippers = client.Database("logistics").Collection("shippers")
}
