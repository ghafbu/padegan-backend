package mongodbpkg

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func NewMongoDB(paramMongoConfig Config) (*mongo.Client, error) {
	// اتصال به MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if paramMongoConfig.Url == "" {
		log.Fatal("MongoDB Error: Set your 'URI' environment variable.")
	}
	thisClient, err := mongo.Connect(options.Client().ApplyURI(paramMongoConfig.Url))
	if err != nil {
		panic(err)
	}
	err = thisClient.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	return thisClient, nil
}
